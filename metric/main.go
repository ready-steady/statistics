// Package metric provides functions for comparing datasets.
package metric

import (
	"math"
)

// MSE computes the mean-square error.
//
// https://en.wikipedia.org/wiki/Mean_squared_error
func MSE(data1, data2 []float64) float64 {
	var sum, Δ float64

	for i := range data1 {
		Δ = data2[i] - data1[i]
		sum += Δ * Δ
	}

	return sum / float64(len(data1))
}

// RMSE computes the root-mean-square error.
//
// https://en.wikipedia.org/wiki/Root-mean-square_deviation
func RMSE(data1, data2 []float64) float64 {
	return math.Sqrt(MSE(data1, data2))
}

// NRMSE computes the normalized root-mean-square error.
//
// https://en.wikipedia.org/wiki/Root-mean-square_deviation#Normalized_root-mean-square_deviation
func NRMSE(observations, predictions []float64) float64 {
	count := len(observations)

	min, max := observations[0], observations[0]
	for i := 1; i < count; i++ {
		if observations[i] < min {
			min = observations[i]
		}
		if observations[i] > max {
			max = observations[i]
		}
	}

	return RMSE(observations, predictions) / (max - min)
}
