package metric

import (
	"math"
	"sort"

	"github.com/ready-steady/linear/metric"
	"github.com/ready-steady/statistics/distribution"
)

// KolmogorovSmirnov computes the Kolmogorov–Smirnov statistic.
//
// https://en.wikipedia.org/wiki/Kolmogorov%E2%80%93Smirnov_test
func KolmogorovSmirnov(data1, data2 []float64) float64 {
	edges := detect(data1, data2)
	return metric.Uniform(distribution.CDF(data1, edges),
		distribution.CDF(data2, edges))
}

// KullbackLeibler computes the Kullback–Leibler divergence.
//
// https://en.wikipedia.org/wiki/Kullback%E2%80%93Leibler_divergence
func KullbackLeibler(pdata, qdata []float64) float64 {
	edges := detect(pdata, qdata)

	pcdf := distribution.CDF(pdata, edges)
	qcdf := distribution.CDF(qdata, edges)

	divergence := 0.0
	for i := range pcdf {
		if pcdf[i] > 0 && qcdf[i] > 0 {
			divergence += pcdf[i] * math.Log(pcdf[i]/qcdf[i])
		}
	}

	return divergence
}

func detect(data1, data2 []float64) []float64 {
	n1, n2 := len(data1), len(data2)
	n := n1 + n2

	edges := make([]float64, n+2)
	edges[0] = math.Inf(-1)
	edges[1] = -edges[0]
	copy(edges[2:], data1)
	copy(edges[2+n1:], data2)

	edges = edges[:unique(edges)]
	sort.Float64s(edges)

	return edges
}

func unique(data []float64) int {
	n := len(data) - 1
	for i := 0; i < n; i++ {
		for j := i + 1; j <= n; j++ {
			if data[i] == data[j] {
				data[j] = data[n]
				n--
				j--
			}
		}
	}
	return n + 1
}
