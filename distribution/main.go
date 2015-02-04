// Package distribution provides tools for working with empirical distributions.
package distribution

// CDF calculates an empirical cummulative distribution function. The
// granularity of the function is specified by a set of edges. For n edges, the
// function returns (n-1) points (see Histogram).
func CDF(data, edges []float64) []float64 {
	bins := Histogram(data, edges)

	total := uint(0)
	for _, count := range bins {
		total += count
	}

	values := make([]float64, len(bins))
	for i := range values {
		values[i] = float64(bins[i]) / float64(total)
		if i > 0 {
			values[i] += values[i-1]
		}
	}

	return values
}

// Histogram counts the number of points that fall into each of the bins
// specified by a set of edges. For n edges, the number of bins is (n-1). The
// left endpoint of a bin is assumed to belong to the bin while the right one is
// assumed to do not.
func Histogram(data []float64, edges []float64) []uint {
	bins := make([]uint, len(edges)-1)

	for _, x := range data {
		if i := find(x, edges); i != -1 {
			bins[i]++
		}
	}

	return bins
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
