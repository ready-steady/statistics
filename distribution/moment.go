package distribution

// Mean computes an estimate of the population mean from a finite sample.
func Mean(data []float64) float64 {
	sum := 0.0

	for _, x := range data {
		sum += x
	}

	return sum / float64(len(data))
}

// Variance computes an estimate of the population variance from a finite
// sample. The estimate is unbiased.
func Variance(data []float64) float64 {
	n, μ := float64(len(data)), Mean(data)

	Σ1, Σ2 := 0.0, 0.0
	for _, x := range data {
		Σ1 = Σ1 + (x-μ)*(x-μ)
		Σ2 = Σ2 + (x - μ)
	}

	return (Σ1 - Σ2*Σ2/n) / (n - 1)
}
