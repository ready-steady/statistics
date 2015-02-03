package test

import (
	"math"
	"sort"

	"github.com/ready-steady/statistics/distribution"
)

// KolmogorovSmirnov performs the Kolmogorov–Smirnov test. The null hypothesis
// is that the data in two data sets are comming from the same continuous
// distribution. The α parameters specifies the significance level.
func KolmogorovSmirnov(points1, points2 []float64, α float64) (bool, float64) {
	const (
		terms = 101
	)

	pc1, pc2 := len(points1), len(points2)
	γ := math.Sqrt(float64(pc1*pc2) / float64(pc1+pc2))

	Δ := computeDistance(points1, points2)

	λ2 := (γ + 0.12 + 0.11/γ) * Δ
	if λ2 < 0 {
		λ2 = 0
	}
	λ2 *= λ2

	p := 0.0

	flip, flop := 1.0, 1.0
	for i := 0; i < terms; i++ {
		p += flip * math.Exp(-2*λ2*flop*flop)
		flip, flop = -flip, flop+1
	}

	p *= 2
	if p < 0 {
		p = 0
	} else if p > 1 {
		p = 1
	}

	return α >= p, p
}

func computeDistance(points1, points2 []float64) float64 {
	edges := computeEdges(points1, points2)

	cdf1 := computeCDF(points1, edges)
	cdf2 := computeCDF(points2, edges)

	var δ, Δ float64

	for i := range cdf1 {
		δ = cdf1[i] - cdf2[i]
		if δ < 0 {
			δ = -δ
		}
		if δ > Δ {
			Δ = δ
		}
	}

	return Δ
}

func computeEdges(points1, points2 []float64) []float64 {
	pc1, pc2 := len(points1), len(points2)
	pc := pc1 + pc2

	edges := make([]float64, pc+2)

	edges[0] = math.Inf(-1)
	copy(edges[1:], points1)
	copy(edges[1+pc1:], points2)
	edges[pc+1] = -edges[0]

	sort.Float64s(edges)

	return edges
}

func computeCDF(points, edges []float64) []float64 {
	bins := distribution.Histogram(points, edges)

	total := uint(0)
	for _, count := range bins {
		total += count
	}

	cdf := make([]float64, len(bins))
	for i := range cdf {
		cdf[i] = float64(bins[i]) / float64(total)
		if i > 0 {
			cdf[i] += cdf[i-1]
		}
	}

	return cdf
}
