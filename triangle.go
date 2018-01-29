package phash

import (
	"context"
	"image"
	"log"
	"math"
	"sort"

	"github.com/azr/phash/cmd"
)

// A Triangle is made of three 2D points.
type Triangle [3]image.Point

func (t *Triangle) A() image.Point { return t[0] }
func (t *Triangle) B() image.Point { return t[1] }
func (t *Triangle) C() image.Point { return t[2] }

// Area of the triangle
func (t *Triangle) Area() int {
	x0 := t.A().X
	x1 := t.B().X
	x2 := t.C().X

	y0 := t.A().Y
	y1 := t.B().Y
	y2 := t.C().Y

	return (x0*(y1-y2) + x1*(y2-y0) + x2*(y0-y1)) / 2
}

// EveryTrianglesOpts is an option for EveryTriangles func
type EveryTrianglesOpts struct {
	Ctx context.Context
	// number of pixels in image
	Pixels int
	// min distance ratio between two points
	// 1 for LowerThreshold would mean a triangle has to
	// contain as much pixel as the image has to be valid
	LowerThresholdRatio, UpperThresholdRatio   float32
	lowerThreshold, upperThreshold             int
	lowerSquareThreshold, upperSquareThreshold int
	MinAreaRatio                               float32 // ratio in pixels an triangle has to be for it to be valid
	minArea                                    int     // in pixel ratio, 1 means everything
}

func (opts *EveryTrianglesOpts) init() {
	opts.lowerThreshold = int(opts.LowerThresholdRatio * float32(opts.Pixels))
	opts.upperThreshold = int(opts.UpperThresholdRatio * float32(opts.Pixels))
	opts.minArea = int(opts.MinAreaRatio * float32(opts.Pixels))
	opts.lowerSquareThreshold, opts.upperSquareThreshold = opts.lowerThreshold*opts.lowerThreshold, opts.upperThreshold*opts.upperThreshold
}

// DistanceInvalid returns true when distance is invalid
//go:nosplit
func (opts *EveryTrianglesOpts) DistanceInvalid(one, two image.Point) bool {
	if true { // so we don't have to sqrt+type change
		distance := squareDistance(one, two)
		return distance < opts.lowerSquareThreshold || distance > opts.upperSquareThreshold
	}
	// legacy code to be sure
	distance := int(math.Sqrt(float64(squareDistance(one, two))))
	return distance < opts.lowerThreshold || distance > opts.upperThreshold
}

// EveryTriangles returns every possible triangle from the points.
func (points Points) EveryTriangles(opts EveryTrianglesOpts) []Triangle {
	opts.init()
	if opts.lowerThreshold == opts.upperThreshold {
		log.Println("EveryTriangles: Identical tresholds, this is not going to work.")
		return nil
	}
	// sort points in order to have cpu
	// do faster using branching
	sort.Sort(points)

	res := []Triangle{}
	for i := 0; i < len(points); i++ {
		tri := points[i:].trianglesToFirst(opts)
		if len(tri) != 0 {
			res = append(res, tri...)
		}
	}
	return res
}

//trianglesToFirst gives you every possible triangles from points[0]
//go:nosplit
func (points Points) trianglesToFirst(opts EveryTrianglesOpts) []Triangle {
	if len(points) < 3 {
		return nil
	}
	X := points[0]
	res := make([]Triangle, 0, len(points))
	for end := 2; end < len(points); end++ {
		Z := points[end]
		if opts.DistanceInvalid(X, Z) {
			continue
		}

		for i := 1; i < end; i++ {
			Y := points[i]
			if opts.DistanceInvalid(X, Y) || opts.DistanceInvalid(Y, Z) {
				continue
			}
			t := Triangle{X, Y, Z}
			area := t.Area()
			if area < opts.minArea {
				continue
			}
			res = append(res, t)
		}
	}
	return res
}

// Debug outputs your image to a triangle.jpg image
// containing only triangle coordinates.
func (t *Triangle) Debug(img image.Image) {
	bounds := t.Bounds()
	subImg := img.(interface {
		SubImage(r image.Rectangle) image.Image
	}).SubImage(bounds)
	cmd.WriteImageToPath(subImg, "square-triangle")
	println("")
}

// Bounds returns a rectangle containing triangle
func (t *Triangle) Bounds() (res image.Rectangle) {
	res.Min.Y, res.Min.X = t[0].Y, t[0].X
	res.Max.Y, res.Max.X = t[0].Y, t[0].X

	for i := 1; i < 3; i++ {
		point := t[i]
		if point.X < res.Min.X {
			res.Min.X = point.X
		}
		if point.X > res.Min.X {
			res.Max.X = point.X
		}
		if point.Y < res.Min.Y {
			res.Min.Y = point.Y
		}
		if point.Y > res.Min.Y {
			res.Max.Y = point.Y
		}
	}
	return res
}
