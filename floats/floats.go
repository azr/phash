package floats

import (
	"math"
)

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

// Round returns the rounded value of val
// When leftover digit is >= to roundOn it is Ceiled
// else Floored
func Round(val, roundOn float64, precision int) float64 {
	var round float64
	pow := math.Pow(10, float64(precision))
	digit := pow * val
	_, div := math.Modf(digit)
	if div >= roundOn {
		round = math.Ceil(digit)
	} else {
		round = math.Floor(digit)
	}
	return round / pow
}
