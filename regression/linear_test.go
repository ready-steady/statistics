package regression

import (
	"testing"

	"github.com/ready-steady/assert"
)

func TestSimpleLinearNew(t *testing.T) {
	x := []float64{
		3.181500e+02,
		3.281500e+02,
		3.381500e+02,
		3.481500e+02,
		3.581500e+02,
		3.681500e+02,
		3.781500e+02,
		3.881500e+02,
		3.981500e+02,
	}
	y := []float64{
		5.459861e-01,
		6.303621e-01,
		7.325854e-01,
		8.549895e-01,
		1.000000e+00,
		1.171091e+00,
		1.373428e+00,
		1.606733e+00,
		1.873658e+00,
	}

	model := NewSimpleLinear(x, y)

	assert.Close(model.α, -4.7845715178610932e+00, 1e-12, t)
	assert.Close(model.β, +1.6395978333333290e-02, 1e-12, t)
}
