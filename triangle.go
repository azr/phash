package phash

import (
	"image"
)

// A Triangle is made of three 2D points.
type Triangle struct {
	A, B, C image.Point
}

// Area of the triangle
func (t *Triangle) Area() int {
	x0 := t.A.X
	x1 := t.B.X
	x2 := t.C.X

	y0 := t.A.Y
	y1 := t.B.Y
	y2 := t.C.Y

	return int(x0*(y1-y2)+x1*(y2-y0)+x2*(y0-y1)) / 2
}

// EveryTrianglesOpts is an option for EveryTriangles func
type EveryTrianglesOpts struct {
	lowerThreshold, upperThreshold float64 // min distance between two points in pixels
	minArea                        int     // in pixels
}

// DistanceInvalid returns true when distance is invalid
func (opts *EveryTrianglesOpts) DistanceInvalid(one, two image.Point) bool {
	distance := distance(one, two)
	return distance < opts.lowerThreshold || distance > opts.upperThreshold
}

// EveryTriangles returns every possible triangle from the points.
// TODO(azr): parallelize.
func (points Points) EveryTriangles(opts EveryTrianglesOpts) []Triangle {
	ts := make([]Triangle, 0, len(points)*len(points)*len(points)) // TODO(azr): make me correct.
	for i := 0; i < len(points); i++ {
		ts = append(ts, points[i:].trianglesToFirst(opts)...)
	}
	return ts
}

//trianglesToFirst gives you every possible triangles from points[0]
func (points Points) trianglesToFirst(opts EveryTrianglesOpts) []Triangle {
	if len(points) < 3 {
		return nil
	}
	X := points[0]
	res := make([]Triangle, 0, len(points)) // TODO(azr): make me correct.
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
			t := Triangle{A: X, B: Y, C: Z}
			if t.Area() < opts.minArea {
				continue
			}
			res = append(res, t)
		}
	}
	return res
}
