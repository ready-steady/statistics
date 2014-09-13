// Package assess provides functions for measuring the difference between the
// values predicted by a model and the values actually observed.
package assess

import (
	"math"
)

// MSE computes the mean-square error.
//
// https://en.wikipedia.org/wiki/Mean_squared_error
func MSE(predicted, observed []float64) float64 {
	var delta, sum float64

	for i := range observed {
		delta = predicted[i] - observed[i]
		sum += delta * delta
	}

	return sum / float64(len(observed))
}

// RMSE computes the root-mean-square error.
//
// https://en.wikipedia.org/wiki/Root-mean-square_deviation
func RMSE(predicted, observed []float64) float64 {
	return math.Sqrt(MSE(predicted, observed))
}

// NRMSE computes the normalized root-mean-square error.
//
// https://en.wikipedia.org/wiki/Root-mean-square_deviation#Normalized_root-mean-square_deviation
func NRMSE(predicted, observed []float64) float64 {
	count := len(observed)
	if count == 0 {
		return 0
	}

	min, max := observed[0], observed[0]
	for i := 1; i < count; i++ {
		if observed[i] < min {
			min = observed[i]
		}
		if observed[i] > max {
			max = observed[i]
		}
	}

	return RMSE(predicted, observed) / (max - min)
}
