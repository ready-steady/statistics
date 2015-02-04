package test

import (
	"math"
	"sort"

	"github.com/ready-steady/statistics/distribution"
)

// KolmogorovSmirnov performs the Kolmogorov–Smirnov test. The null hypothesis
// is that the data in two data sets are comming from the same continuous
// distribution. The α parameters specifies the significance level.
func KolmogorovSmirnov(data1, data2 []float64, α float64) (bool, float64) {
	const (
		terms = 101
	)

	pc1, pc2 := len(data1), len(data2)
	γ := math.Sqrt(float64(pc1*pc2) / float64(pc1+pc2))

	edges := computeEdges(data1, data2)
	Δ := computeInfNorm(computeCDF(data1, edges), computeCDF(data2, edges))

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

func computeEdges(data1, data2 []float64) []float64 {
	pc1, pc2 := len(data1), len(data2)
	pc := pc1 + pc2

	edges := make([]float64, pc+2)

	edges[0] = math.Inf(-1)
	copy(edges[1:], data1)
	copy(edges[1+pc1:], data2)
	edges[pc+1] = -edges[0]

	sort.Float64s(edges)

	return edges
}

func computeCDF(data, edges []float64) []float64 {
	bins := distribution.Histogram(data, edges)

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

func computeInfNorm(data1, data2 []float64) float64 {
	var δ, Δ float64

	for i := range data1 {
		δ = data1[i] - data2[i]
		if δ < 0 {
			δ = -δ
		}
		if δ > Δ {
			Δ = δ
		}
	}

	return Δ
}
