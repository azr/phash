package phash

import (
	"context"
	"image"
	"log"
	"math"
	"sort"
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

	return int(x0*(y1-y2)+x1*(y2-y0)+x2*(y0-y1)) / 2
}

// EveryTrianglesOpts is an option for EveryTriangles func
type EveryTrianglesOpts struct {
	Ctx context.Context
	// min distance between two points in pixels
	// 150 / 500 respecively are good values
	LowerThreshold, UpperThreshold             float64
	lowerSquareThreshold, upperSquareThreshold int
	MinArea                                    int // in pixels, 50 is a good value
}

func (opts *EveryTrianglesOpts) init() {
	opts.lowerSquareThreshold, opts.upperSquareThreshold = int(math.Pow(opts.LowerThreshold, 2)), int(math.Pow(opts.UpperThreshold, 2))
}

// DistanceInvalid returns true when distance is invalid
//go:nosplit
func (opts *EveryTrianglesOpts) DistanceInvalid(one, two image.Point) bool {
	distance := squareDistance(one, two)
	return distance < opts.lowerSquareThreshold || distance > opts.upperSquareThreshold
}

// EveryTriangles returns every possible triangle from the points.
// TODO(azr): parallelize.
func (points Points) EveryTriangles(opts EveryTrianglesOpts) []Triangle {
	if opts.LowerThreshold == 0 {
		opts.LowerThreshold = 150
	}
	if opts.UpperThreshold == 0 {
		opts.UpperThreshold = 500
	}
	if opts.MinArea == 0 {
		opts.MinArea = 50
	}
	if opts.LowerThreshold == opts.UpperThreshold {
		log.Println("EveryTriangles: Identical tresholds, this is not going to work.")
		return nil
	}
	opts.init()
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
	log.Printf("EveryTriangles found %d triangles", len(res))
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
			if t.Area() < opts.MinArea {
				continue
			}
			res = append(res, t)
		}
	}
	return res
}
