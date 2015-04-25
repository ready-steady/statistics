package metric

import (
	"math"
	"math/rand"
	"testing"

	"github.com/ready-steady/assert"
)

func TestDetect(t *testing.T) {
	data1 := []float64{1, math.Inf(1), 2, 0}
	data2 := []float64{0, 2, 4, 1, 1, 1, 4}
	edges := []float64{math.Inf(-1), 0, 1, 2, 4, math.Inf(1)}
	assert.Equal(detect(data1, data2), edges, t)
}

func TestSortUnique(t *testing.T) {
	data := []float64{1, math.Inf(1), 2, 0, 0, 2, 4, 1, 1, 1, 4}
	data = data[:sortUnique(data)]
	assert.Equal(data, []float64{0, 1, 2, 4, math.Inf(1)}, t)
}

func BenchmarkSortUnique(b *testing.B) {
	for i := 0; i < b.N; i++ {
		sortUnique(random(1e6, 10))
	}
}

func random(count, max uint) []float64 {
	data := make([]float64, count)
	for i := range data {
		data[i] = float64(uint(float64(max) * rand.Float64()))
	}
	return data
}
