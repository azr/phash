package phash

import (
	"image"

	"github.com/anthonynsimon/bild/effect"
)

// Points is a slice of points.
type Points []image.Point

// FindKeypoints returns a list of points that are key points
// it first does an edge detection then appends
// the 'set' pixels coordinates into an array.
func FindKeypoints(src image.Image) Points {
	edgySrc := effect.EdgeDetection(src, 1) // TODO (azr): Edge Detection should return a image.Grey or Alpha for efficiency
	bounds := edgySrc.Bounds()
	points := make([]image.Point, 0, len(edgySrc.Pix)/2) // this is probably a premature optimization
	for x := 0; x < bounds.Dx(); x++ {
		for y := 0; y < bounds.Dy(); y++ {
			pix := edgySrc.RGBAAt(x, y)
			if pix.A != 0 {
				// image is greyscale &
				// there is an alpha so this point is interesting :)
				points = append(points, image.Point{X: y, Y: y})
			}
		}
	}
	return points
}
