package phash

import (
	"image"

	"github.com/azr/gift"
	"github.com/azr/phash/geometry"
)

// FindKeypoints returns a list of points that are interesting/key points
// It does that by detecting corners using cornerdectect.Fast9
func FindKeypoints(src image.Image) geometry.Points {
	g := gift.New(
		gift.Grayscale(),
	)

	simplifiedImage := image.NewGray(g.Bounds(src.Bounds()))
	g.Draw(simplifiedImage, src)
	points := geometry.DetectCorners(simplifiedImage)
	return points
}
