// Package corr provides functions for working with correlation coefficients
// and correlation/covariance matrices.
package corr

import (
	"math"

	"github.com/go-math/stat/decomp"
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

// Decompose computes an m-by-n matrix M for an m-by-m covariance matrix Σ such
// that, for an n-element vector X with uncorrelated components, M * X is an
// m-element vector whose components are correlated according to Σ. The
// function reduces the number of dimensions from m to n such that a certain
// portion of the variance is preserved, which is controlled by λ ∈ (0, 1].
func Decompose(Σ []float64, m uint32, λ float64) ([]float64, uint32, error) {
	U, Λ, err := decomp.CovPCA(Σ, m)
	if err != nil {
		return nil, 0, err
	}

	n := m

	// NOTE: Λ is in descending order and non-negative.
	var cum, sum float64
	for i := uint32(0); i < m; i++ {
		sum += Λ[i]
	}
	for i := uint32(0); i < m; i++ {
		cum += Λ[i]
		if cum/sum >= λ {
			n = i + 1
			break
		}
	}

	for i := uint32(0); i < n; i++ {
		coef := math.Sqrt(Λ[i])
		for j := uint32(0); j < m; j++ {
			U[i*m+j] *= coef
		}
	}

	return U[:m*n], n, nil
}
