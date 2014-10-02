package radon

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func roundingFactor(x float64) float64 {
	if (x) >= 0 {
		return 0.5
	}
	return -0.5
}
