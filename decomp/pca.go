package decomp

import (
	"github.com/go-math/linal/decomp"
)

// CovPCA performs principal component analysis of an m-by-m covariance matrix.
func CovPCA(Σ []float64, m uint32) ([]float64, []float64, error) {
	U := make([]float64, m*m)
	Λ := make([]float64, m)

	if err := decomp.SymEigen(Σ, U, Λ, m); err != nil {
		return nil, nil, err
	}

	// NOTE: The eigenvalues in Λ are in ascending order. Reverse!
	for i, j := uint32(0), m-1; i < j; i, j = i+1, j-1 {
		Λ[i], Λ[j] = Λ[j], Λ[i]
		for k := uint32(0); k < m; k++ {
			U[i*m+k], U[j*m+k] = U[j*m+k], U[i*m+k]
		}
	}

	return U, Λ, nil
}
