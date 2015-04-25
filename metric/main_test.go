package metric

import (
	"math"
	"testing"

	"github.com/ready-steady/assert"
)

func TestDetect(t *testing.T) {
	data1 := []float64{1, math.Inf(1), 2, 0}
	data2 := []float64{0, 2, 4, 1, 1, 1, 4}
	edges := []float64{math.Inf(-1), 0, 1, 2, 4, math.Inf(1)}

	assert.Equal(detect(data1, data2), edges, t)
}
