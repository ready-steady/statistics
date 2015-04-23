package metric

import (
	"math"
	"sort"
	"testing"

	"github.com/ready-steady/assert"
)

func TestDetect(t *testing.T) {
	data1 := []float64{1, math.Inf(1), 2, 0}
	data2 := []float64{0, 2, 4, 1, 1, 1, 4}

	edges := detect(data1, data2)

	assert.Equal(edges, []float64{math.Inf(-1), 0, 1, 2, 4, math.Inf(1)}, t)
}

func TestUnique(t *testing.T) {
	data := []float64{0, 1, 2, 3, 1, 2, 1, 4, 5, 0}

	data = data[:unique(data)]
	sort.Float64s(data)

	assert.Equal(data, []float64{0, 1, 2, 3, 4, 5}, t)
}
