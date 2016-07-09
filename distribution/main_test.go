package distribution

import (
	"testing"

	"github.com/ready-steady/assert"
)

func TestHistogram(t *testing.T) {
	data := []float64{
		8.1472368639317894e-01, 9.0579193707561922e-01,
		1.2698681629350606e-01, 9.1337585613901939e-01,
		6.3235924622540951e-01, 9.7540404999409525e-02,
		2.7849821886704840e-01, 5.4688151920498385e-01,
		9.5750683543429760e-01, 9.6488853519927653e-01,
		1.5761308167754828e-01, 9.7059278176061570e-01,
		9.5716694824294557e-01, 4.8537564872284122e-01,
		8.0028046888880011e-01, 1.4188633862721534e-01,
		4.2176128262627499e-01, 9.1573552518906709e-01,
		7.9220732955955442e-01, 9.5949242639290300e-01,
	}
	edges := []float64{0.0, 0.1, 0.2, 0.3, 0.4, 0.5, 0.6, 0.7, 0.8, 0.9, 1.0}
	bins := []uint{1, 3, 1, 0, 2, 1, 1, 1, 2, 8}

	result, _ := Histogram(data, edges)
	assert.Equal(bins, result, t)
}

func TestEdges(t *testing.T) {
	data1 := []float64{1.0, infinity, 2.0, 0.0}
	data2 := []float64{0.0, 2.0, 4.0, 1.0, 1.0, 1.0, 4.0}
	edges := []float64{-infinity, 0.0, 1.0, 2.0, 4.0, infinity}

	assert.Equal(Edges(data1, data2), edges, t)
}

func TestFind(t *testing.T) {
	edges := []float64{-2.0, -1.0, 0.0, 1.0, 2.0}

	cases := []struct {
		x float64
		i int
	}{
		{-10.0, -1.0},
		{-2.0, 0.0},
		{-1.5, 0.0},
		{-1.0, 1.0},
		{-0.5, 1.0},
		{0.0, 2.0},
		{0.5, 2.0},
		{1.0, 3.0},
		{1.5, 3.0},
		{2.0, -1.0},
		{10.0, -1.0},
	}

	for _, c := range cases {
		assert.Equal(find(c.x, edges), c.i, t)
	}
}
