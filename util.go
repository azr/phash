package phash

import (
	"image/color"
)

// ColorToGreyScaleFloat64 reduces rgb
// to a greyscale approximation.
func ColorToGreyScaleFloat64(c color.Color) float64 {
	r, g, b, _ := c.RGBA()
	return 0.299*float64(r) +
		0.587*float64(g) +
		0.114*float64(b)
}
