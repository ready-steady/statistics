package regression

// SimpleLinear is a simple linear regression model.
//
// https://en.wikipedia.org/wiki/Simple_linear_regression
type SimpleLinear struct {
	α float64
	β float64
}

// NewSimpleLinear fits a simple linear regression model.
func NewSimpleLinear(x, y []float64) *SimpleLinear {
	n := uint(len(x))

	μx, μy, μxy, μx2 := 0.0, 0.0, 0.0, 0.0
	for i := uint(0); i < n; i++ {
		μx += x[i]
		μy += y[i]
		μxy += x[i] * y[i]
		μx2 += x[i] * x[i]
	}
	μx, μy, μxy, μx2 = μx/float64(n), μy/float64(n), μxy/float64(n), μx2/float64(n)

	β := (μxy - μx*μy) / (μx2 - μx*μx)
	α := μy - β*μx

	return &SimpleLinear{α, β}
}

// Compute evaluate the model at a point.
func (self *SimpleLinear) Compute(x float64) float64 {
	return self.α + self.β*x
}
