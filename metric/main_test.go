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

	edges := detect(data1, data2)

	assert.Equal(edges, []float64{math.Inf(-1), 0, 1, 2, 4, math.Inf(1)}, t)
}

func TestUnique(t *testing.T) {
	data := []float64{0, 0, 0, 1, 1, 1, 1, 2, 2, 2, 4, 5, 5}
	assert.Equal(data[:unique(data)], []float64{0, 1, 2, 4, 5}, t)
}

func BenchmarkDetect(b *testing.B) {
	rand.Seed(0)
	data1 := random(1e6, 10)
	data2 := random(1e6, 10)
	for i := 0; i < b.N; i++ {
		detect(data1, data2)
	}
}

func random(count, max uint) []float64 {
	data := make([]float64, count)
	for i := range data {
		data[i] = float64(uint(float64(max) * rand.Float64()))
	}
	return data
}
