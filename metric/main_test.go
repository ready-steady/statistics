package metric

import (
	"testing"

	"github.com/ready-steady/assert"
)

func TestDetect(t *testing.T) {
	data1 := []float64{1.0, infinity, 2.0, 0.0}
	data2 := []float64{0.0, 2.0, 4.0, 1.0, 1.0, 1.0, 4.0}
	edges := []float64{-infinity, 0.0, 1.0, 2.0, 4.0, infinity}

	assert.Equal(detect(data1, data2), edges, t)
}
