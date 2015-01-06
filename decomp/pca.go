package decomp

import (
	"errors"

	"github.com/ready-steady/linal/decomp"
)

// CovPCA performs principal component analysis on an m-by-m covariance matrix.
// The principal components U and their variances Λ are returned in descending
// order of Λ.
func CovPCA(Σ []float64, m uint32) (U []float64, Λ []float64, err error) {
	U = make([]float64, m*m)
	Λ = make([]float64, m)

	if err = decomp.SymEig(Σ, U, Λ, m); err != nil {
		return nil, nil, err
	}

	for i := uint32(0); i < m; i++ {
		if Λ[i] < 0 {
			return nil, nil, errors.New("the matrix is not positive semidefinite")
		}
	}

	// NOTE: Λ is in ascending order. Reverse!
	for i, j := uint32(0), m-1; i < j; i, j = i+1, j-1 {
		Λ[i], Λ[j] = Λ[j], Λ[i]
		for k := uint32(0); k < m; k++ {
			U[i*m+k], U[j*m+k] = U[j*m+k], U[i*m+k]
		}
	}

	return
}
