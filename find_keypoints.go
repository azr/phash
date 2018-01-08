package phash

import (
	"image"
	"log"

	"github.com/anthonynsimon/bild/effect"
	"github.com/azr/phash/cmd"
)

// Points is a slice of points.
type Points []image.Point

func (p Points) Len() int           { return len(p) }
func (p Points) Less(i, j int) bool { return p[i].X < p[j].X }
func (p Points) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

// FindKeypoints returns a list of points that are key points
// it first does an edge detection then appends
// the 'set' pixels coordinates into an array.
// treshold is minimum acceptable color value
// given by ColorToGreyScaleFloat64, any value under
// treshold is ignored.
// A good value for threshold is 3000
func FindKeypoints(src image.Image, treshold float64) Points {
	edgySrc := effect.EdgeDetection(src, 1)
	bounds := edgySrc.Bounds()
	points := make([]image.Point, 0, len(edgySrc.Pix)/2) // this is probably a premature optimization
	v := 0.0
	nvs := 0
	for x := 0; x < bounds.Dx(); x++ {
		for y := 0; y < bounds.Dy(); y++ {
			v += ColorToGreyScaleFloat64(edgySrc.At(x, y))
			nvs++
			if ColorToGreyScaleFloat64(edgySrc.At(x, y)) > treshold {
				// image is greyscale &
				// there is an alpha so this point is interesting :)
				points = append(points, image.Point{X: y, Y: y})
			}
		}
	}
	log.Printf("FindKeypoints: v: %f, nvs: %d, v/nvs: %fm kept: %d", v, nvs, v/float64(nvs), len(points))
	return points
}
