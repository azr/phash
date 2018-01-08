package phash

import (
	"image"
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

// squareDistance gives you distance^2
// you'd have to Srqt it for it to be correct
func squareDistance(one, two image.Point) int {
	dx := one.X - two.X
	dy := one.Y - two.Y
	return (dx * dx) + (dy * dy)
}

func min(a, b int) int {
	if b < a {
		return b
	}
	return a
}

func max(a, b int) int {
	if b > a {
		return b
	}
	return a
}
