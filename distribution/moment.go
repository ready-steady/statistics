package distribution

// Expectation computes an estimate of the population mean from a finite sample.
func Expectation(data []float64) float64 {
	Σ := 0.0
	for _, x := range data {
		Σ += x
	}
	return Σ / float64(len(data))
}

// Variance computes an estimate of the population variance from a finite
// sample. The estimate is unbiased. The computation is based on the
// compensated-summation version of the two-pass algorithm.
//
// https://en.wikipedia.org/wiki/Algorithms_for_calculating_variance#Two-pass_algorithm
func Variance(data []float64) float64 {
	n, μ := float64(len(data)), Expectation(data)
	Σ1, Σ2 := 0.0, 0.0
	for _, x := range data {
		Δ := x - μ
		Σ1 += Δ
		Σ2 += Δ * Δ
	}
	return (Σ2 - Σ1*Σ1/n) / (n - 1.0)
}
