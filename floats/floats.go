package floats

import ()

func sum(s []float64, pos int) float64 {
	if pos == 0 {
		return s[0]
	}
	return s[pos] + sum(s, pos-1)
}

// Sum returns the maximum value of a slice recursively
func Sum(s []float64) float64 {
	len := len(s)
	switch len {
	case 0:
		return 0
	default:
		return sum(s, len-1)
	}
}
