// Package metric provides functions for comparing datasets.
package metric

import (
	"math"
)

// MSE computes the mean-square error.
//
// https://en.wikipedia.org/wiki/Mean_squared_error
func MSE(y, yhat []float64) float64 {
	var sum, Δ float64

	for i := range y {
		Δ = yhat[i] - y[i]
		sum += Δ * Δ
	}

	return sum / float64(len(y))
}

// RMSE computes the root-mean-square error.
//
// https://en.wikipedia.org/wiki/Root-mean-square_deviation
func RMSE(y, yhat []float64) float64 {
	return math.Sqrt(MSE(y, yhat))
}

// NRMSE computes the normalized root-mean-square error.
//
// https://en.wikipedia.org/wiki/Root-mean-square_deviation#Normalized_root-mean-square_deviation
func NRMSE(y, yhat []float64) float64 {
	count := len(y)

	min, max := y[0], y[0]
	for i := 1; i < count; i++ {
		if y[i] < min {
			min = y[i]
		}
		if y[i] > max {
			max = y[i]
		}
	}

	return RMSE(y, yhat) / (max - min)
}
