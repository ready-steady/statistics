// Package metric provides tools for comparing datasets.
package metric

import (
	"math"

	"github.com/ready-steady/linear/metric"
	"github.com/ready-steady/statistics/distribution"
)

// KolmogorovSmirnov computes the Kolmogorov–Smirnov statistic for two samples.
//
// https://en.wikipedia.org/wiki/Kolmogorov%E2%80%93Smirnov_test
func KolmogorovSmirnov(data1, data2 []float64) float64 {
	edges := distribution.Edges(data1, data2)
	cdf1 := distribution.CDF(data1, edges)
	cdf2 := distribution.CDF(data2, edges)
	return metric.Uniform(cdf1, cdf2)
}

// KullbackLeibler computes the Kullback–Leibler divergence of q from p where p
// and q are two discrete probability distributions. The distribution p is
// assumed to be absolutely continuous with respect the distribution q, that is,
// q[i] = 0 implies that p[i] = 0.
//
// https://en.wikipedia.org/wiki/Kullback%E2%80%93Leibler_divergence
func KullbackLeibler(p, q []float64) float64 {
	divergence := 0.0
	for i := range p {
		if p[i] > 0.0 {
			divergence += p[i] * math.Log(p[i]/q[i])
		}
	}
	return divergence
}

// MSE computes the mean-square error.
//
// https://en.wikipedia.org/wiki/Mean_squared_error
func MSE(y, yhat []float64) float64 {
	Σ := 0.0
	for i := range y {
		ε := yhat[i] - y[i]
		Σ += ε * ε
	}
	return Σ / float64(len(y))
}

// MSPE computes the mean-square-percentage error.
func MSPE(y, yhat []float64) float64 {
	Σ := 0.0
	for i := range y {
		ε := (yhat[i] - y[i]) / y[i]
		Σ += ε * ε
	}
	return Σ / float64(len(y))
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

// RMSE computes the root-mean-square error.
//
// https://en.wikipedia.org/wiki/Root-mean-square_deviation
func RMSE(y, yhat []float64) float64 {
	return math.Sqrt(MSE(y, yhat))
}

// RMSPE computes the root-mean-square-percentage error.
func RMSPE(y, yhat []float64) float64 {
	return math.Sqrt(MSPE(y, yhat))
}
