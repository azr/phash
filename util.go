package phash

import (
	"image"
	"math"
)

func distance(one, two image.Point) float64 {
	return math.Sqrt(math.Pow(float64(one.X-two.X), 2.0) + math.Pow(float64(one.Y-two.Y), 2.0))
}
