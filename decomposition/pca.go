package decomposition

import (
	"errors"

	"github.com/ready-steady/linear/matrix"
)

// CovPCA performs principal component analysis on an m-by-m covariance matrix
// Σ. The principal components U and their variances Λ are returned in
// descending order of the variances. By definition, the variances should be
// nonnegative. Due to finite-precision arithmetics, however, some close-to-zero
// variances might turn out to be negative. If the absolute value of a negative
// variance is smaller than the tolerance ε, the function nullifies that
// variance and proceeds without any errors; otherwise, an error is returned.
func CovPCA(Σ []float64, m uint, ε float64) (U []float64, Λ []float64, err error) {
	U = make([]float64, m*m)
	Λ = make([]float64, m)

	if err = matrix.SymmetricEigen(Σ, U, Λ, m); err != nil {
		return nil, nil, err
	}

	for i := uint(0); i < m; i++ {
		if Λ[i] < 0 {
			if -Λ[i] < ε {
				Λ[i] = 0.0
			} else {
				return nil, nil, errors.New("the matrix should be positive semidefinite")
			}
		}
	}

	// NOTE: Λ is in ascending order. Reverse!
	for i, j := uint(0), m-1; i < j; i, j = i+1, j-1 {
		Λ[i], Λ[j] = Λ[j], Λ[i]
		for k := uint(0); k < m; k++ {
			U[i*m+k], U[j*m+k] = U[j*m+k], U[i*m+k]
		}
	}

	return
}
