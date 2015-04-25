// Package test provides statistical tests.
package test

import (
	"math"

	"github.com/ready-steady/statistics/metric"
)

// KolmogorovSmirnov performs the two-sample Kolmogorov–Smirnov test. The null
// hypothesis is that the two datasets are coming from the same continuous
// distribution. The α parameter specifies the significance level. If the test
// rejects the null hypothesis, the function returns true; otherwise, false is
// returned. The second and third outputs of the function are the p-value and
// Kolmogorov–Smirnov statistic of the test, respectively.
//
// https://en.wikipedia.org/wiki/Kolmogorov%E2%80%93Smirnov_test
func KolmogorovSmirnov(data1, data2 []float64, α float64) (bool, float64, float64) {
	const (
		terms = 101
	)

	statistic := metric.KolmogorovSmirnov(data1, data2)

	// M. Stephens. Use of the Kolmogorov–Smirnov, Cramer-Von Mises and Related
	// Statistics Without Extensive Tables. Journal of the Royal Statistical
	// Society. Series B (Methodological), vol. 32, no. 1 (1970), pp. 115–122.
	//
	// http://www.jstor.org/stable/2984408
	n1, n2 := len(data1), len(data2)
	γ := math.Sqrt(float64(n1*n2) / float64(n1+n2))
	λ := (γ + 0.12 + 0.11/γ) * statistic

	// Kolmogorov distribution
	//
	// https://en.wikipedia.org/wiki/Kolmogorov%E2%80%93Smirnov_test#Kolmogorov_distribution
	pvalue, sign, k := 0.0, 1.0, 1.0
	for i := 0; i < terms; i++ {
		pvalue += sign * math.Exp(-2*λ*λ*k*k)
		sign, k = -sign, k+1
	}
	pvalue *= 2
	if pvalue < 0 {
		pvalue = 0
	} else if pvalue > 1 {
		pvalue = 1
	}

	return α >= pvalue, pvalue, statistic
}
