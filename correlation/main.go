// Package correlation provides tools for working with correlation coefficients
// and correlation matrices.
package correlation

import (
	"math"

	"github.com/ready-steady/statistics/decomposition"
)

// Decompose computes an m-by-n matrix C and an n-by-m matrix D given an m-by-m
// covariance matrix Σ such that (a) for an n-element vector Z with uncorrelated
// components, C * X is an m-element vector whose components are correlated
// according to Σ, and (b) for an m-element vector X with correlated components
// according to Σ, D * X is an n-element vector whose components are
// uncorrelated.
//
// The function reduces the number of dimensions from m to n such that a certain
// portion of the variance is preserved, which is controlled by λ ∈ (0, 1]. For
// the meaning of the ε parameter, refer to decomposition.PCA.
func Decompose(Σ []float64, m uint, λ, ε float64) ([]float64, []float64, uint, error) {
	U, Λ, err := decomposition.CovPCA(Σ, m, ε)
	if err != nil {
		return nil, nil, 0, err
	}

	C, n := U, m

	var cum, sum float64
	for i := uint(0); i < m; i++ {
		sum += Λ[i]
	}
	for i := uint(0); i < m; i++ {
		cum += Λ[i]
		if cum/sum >= λ {
			n = i + 1
			break
		}
	}

	D := make([]float64, n*m)

	for i := uint(0); i < n; i++ {
		σ := math.Sqrt(Λ[i])
		for j := uint(0); j < m; j++ {
			ρ := C[i*m+j]
			C[i*m+j] = ρ * σ
			D[j*n+i] = ρ / σ
		}
	}

	return C[:m*n], D, n, nil
}

// KendallPearson converts the Kendall τ rank correlation coefficient into
// the Pearson correlation coefficient.
//
// https://en.wikipedia.org/wiki/Kendall_tau_rank_correlation_coefficient
func KendallPearson(τ []float64) []float64 {
	r := make([]float64, len(τ))
	for i := range r {
		r[i] = math.Sin(math.Pi * τ[i] / 2.0)
	}
	return r
}

// SpearmanPearson converts Spearman’s rank correlation coefficient into the
// Pearson correlation coefficient.
//
// https://en.wikipedia.org/wiki/Spearman%27s_rank_correlation_coefficient
func SpearmanPearson(ρ []float64) []float64 {
	r := make([]float64, len(ρ))
	for i := range r {
		r[i] = 2.0 * math.Sin(math.Pi*ρ[i]/6.0)
	}
	return r
}
