package geometry

import "image"

// Min between a and b
func Min(a, b int) int {
	if b < a {
		return b
	}
	return a
}

// SquareDistance gives you distance^2 between two image.Point
// you'd have to Srqt it for it to be the correct distance.
func SquareDistance(one, two image.Point) int {
	dx := one.X - two.X
	dy := one.Y - two.Y
	return (dx * dx) + (dy * dy)
}
