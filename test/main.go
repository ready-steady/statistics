package test

import (
	"math"
	"sort"

	"github.com/ready-steady/linear/metric"
	"github.com/ready-steady/statistics/distribution"
)

// KolmogorovSmirnov performs the Kolmogorov–Smirnov test. The null hypothesis
// is that the data in two data sets are comming from the same continuous
// distribution. The test tries to reject the null hypothesis, and the function
// returns true if the null hypothesis has been rejected. The α parameters
// specifies the significance level.
func KolmogorovSmirnov(data1, data2 []float64, α float64) (bool, float64) {
	const (
		terms = 101
	)

	edges := computeEdges(data1, data2)
	Δ := metric.Uniform(distribution.CDF(data1, edges), distribution.CDF(data2, edges))

	// M. Stephens. Use of the Kolmogorov–Smirnov, Cramer-Von Mises and Related
	// Statistics Without Extensive Tables. Journal of the Royal Statistical
	// Society. Series B (Methodological), vol. 32, no. 1 (1970), pp. 115–122.
	//
	// http://www.jstor.org/stable/2984408
	pc1, pc2 := len(data1), len(data2)
	γ := math.Sqrt(float64(pc1*pc2) / float64(pc1+pc2))
	Δ = (γ + 0.12 + 0.11/γ) * Δ

	// Kolmogorov distribution
	//
	// https://en.wikipedia.org/wiki/Kolmogorov–Smirnov_test#Kolmogorov_distribution
	q, sign, k := 0.0, 1.0, 1.0
	for i := 0; i < terms; i++ {
		q += sign * math.Exp(-2*Δ*Δ*k*k)
		sign, k = -sign, k+1
	}
	q *= 2
	if q < 0 {
		q = 0
	} else if q > 1 {
		q = 1
	}

	return α >= q, q
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
