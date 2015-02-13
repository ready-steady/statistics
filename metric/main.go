// Package metric provides functions for measuring the difference between the
// values predicted by a model and the values actually observed.
package metric

import (
	"math"
)

// MSE computes the mean-square error.
//
// https://en.wikipedia.org/wiki/Mean_squared_error
func MSE(predictions, observations []float64) float64 {
	var sum, Δ float64

	for i := range observations {
		Δ = predictions[i] - observations[i]
		sum += Δ * Δ
	}

	return sum / float64(len(observations))
}

// RMSE computes the root-mean-square error.
//
// https://en.wikipedia.org/wiki/Root-mean-square_deviation
func RMSE(predictions, observations []float64) float64 {
	return math.Sqrt(MSE(predictions, observations))
}

// NRMSE computes the normalized root-mean-square error.
//
// https://en.wikipedia.org/wiki/Root-mean-square_deviation#Normalized_root-mean-square_deviation
func NRMSE(predictions, observations []float64) float64 {
	count := len(observations)
	if count == 0 {
		return 0
	}

	min, max := observations[0], observations[0]
	for i := 1; i < count; i++ {
		if observations[i] < min {
			min = observations[i]
		}
		if observations[i] > max {
			max = observations[i]
		}
	}

	return RMSE(predictions, observations) / (max - min)
}
