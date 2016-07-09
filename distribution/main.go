// Package distribution provides tools for working with empirical distributions.
package distribution

import (
	"math"

	"github.com/ready-steady/sort"
)

var (
	infinity = math.Inf(1.0)
)

// CDF calculates an empirical cumulative distribution function. The granularity
// of the function is specified by a set of edges; see Histogram.
func CDF(data, edges []float64) (values []float64) {
	bins, _ := Histogram(data, edges)
	total := uint(0)
	for i, count := range bins {
		total += count
		bins[i] = total
	}
	values = make([]float64, len(bins))
	for i := range values {
		values[i] = float64(bins[i]) / float64(total)
	}
	return
}

// Edges returns sorted unique elements of a number of data sets, ensuring that
// the first and last elements are -∞ and +∞, respectively.
func Edges(data ...[]float64) []float64 {
	count := len(data)
	total := 0
	for i := 0; i < count; i++ {
		total += len(data[i])
	}
	edges := make([]float64, 1+total+1)
	edges[0] = -infinity
	for i, j := 0, 1; i < count; i++ {
		copy(edges[j:], data[i])
		j += len(data[i])
	}
	edges[1+total] = infinity
	return sort.Unique(edges)
}

// Histogram counts the number of points that fall into each of the bins
// specified by a set of edges. For n edges, the number of bins is (n-1). The
// left endpoint of a bin is assumed to belong to the bin while the right one is
// assumed to do not.
func Histogram(data []float64, edges []float64) (bins []uint, total uint) {
	bins = make([]uint, len(edges)-1)
	for _, x := range data {
		if i := find(x, edges); i != -1 {
			bins[i]++
			total++
		}
	}
	return
}

// PDF calculates an empirical probability density function. The granularity of
// the function is specified by a set of edges; see Histogram.
func PDF(data, edges []float64) (values []float64) {
	bins, total := Histogram(data, edges)
	values = make([]float64, len(bins))
	for i := range values {
		values[i] = float64(bins[i]) / float64(total)
	}
	return
}

func find(x float64, edges []float64) int {
	lower, upper := 0, len(edges)-1
	if x < edges[lower] {
		return -1
	}
	if x >= edges[upper] {
		return -1
	}
	for upper-lower > 1 {
		middle := (upper + lower) / 2

		if x >= edges[middle] {
			lower = middle
		} else {
			upper = middle
		}
	}
	return lower
}
