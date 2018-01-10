package phash

import (
	"image"
	"log"

	"github.com/anthonynsimon/bild/effect"
	"github.com/disintegration/gift"
)

// Points is a slice of points.
type Points []image.Point

func (p Points) Len() int           { return len(p) }
func (p Points) Less(i, j int) bool { return p[i].X < p[j].X }
func (p Points) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

// FindKeypoints returns a list of points that are key points
// it first does an edge detection then appends
// the 'set' pixels coordinates into an array.
func FindKeypoints(src image.Image) Points {
	g := gift.New(
		gift.GaussianBlur(2.1),
		gift.Grayscale(),
		gift.Threshold(50),
	)
	simplifiedImage := image.NewGray(g.Bounds(src.Bounds()))
	g.Draw(simplifiedImage, src)
	edgySrc := effect.EdgeDetection(simplifiedImage, 1)
	bounds := edgySrc.Bounds()
	// TODO:
	// Find contours
	// for each contours that have an area bigger than 200
	// grab momens & keep cX, cY = int(M["m10"] / M["m00"]), int(M["m01"] / M["m00"]) => points
	// grab max curvature points ( steep angles )									  => points
	points := make([]image.Point, 0, len(edgySrc.Pix)/2) // this is probably a premature optimization
	for x := 0; x < bounds.Dx(); x++ {
		for y := 0; y < bounds.Dy(); y++ {
			if _, _, b, _ := edgySrc.At(x, y).RGBA(); b > 0 {
				points = append(points, image.Point{X: x, Y: y})
			}
		}
	}
	log.Printf("FindKeypoints kept %d points", len(points))
	return points
}
