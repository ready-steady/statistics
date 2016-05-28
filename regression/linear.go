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
	α := 0.0
	β := 0.0
	return &SimpleLinear{α, β}
}

// Compute evaluate the model at a point.
func (self *SimpleLinear) Compute(x float64) float64 {
	return self.α*x + self.β
}
