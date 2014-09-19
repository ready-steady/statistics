// Package corr provides functions for working with correlation coefficients.
package corr

import (
	"math"
)

// SpearmanPearson converts Spearman’s rank correlation coefficient into the
// Pearson correlation coefficient.
//
// https://en.wikipedia.org/wiki/Spearman%27s_rank_correlation_coefficient
func SpearmanPearson(ρ []float64) []float64 {
	r := make([]float64, len(ρ))

	for i := range r {
		r[i] = 2 * math.Sin(math.Pi*ρ[i]/6)
	}

	return r
}

// KendallPearson converts the Kendall τ rank correlation coefficient into
// the Pearson correlation coefficient.
//
// https://en.wikipedia.org/wiki/Kendall_tau_rank_correlation_coefficient
func KendallPearson(τ []float64) []float64 {
	r := make([]float64, len(τ))

	for i := range r {
		r[i] = math.Sin(math.Pi * τ[i] / 2)
	}

	return r
}
