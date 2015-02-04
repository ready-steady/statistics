// Package distribution provides tools for working with empirical distributions.
package distribution

// Histogram counts the number of points that fall into each of the bins
// specified by the given edges. For n edges, the number of bins is (n-1). The
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
