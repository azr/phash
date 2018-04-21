package triangle

import (
	"image"
	"log"
	"sort"

	"github.com/azr/phash/geometry"
)

// PossibilititesOpts is an option for EveryTriangles func
type PossibilititesOpts struct {
	// min distance ratio between two points
	LowerThreshold, UpperThreshold             int
	lowerSquareThreshold, upperSquareThreshold int // pow 2 of thresholds

	MinArea int // in pixel
}

func (opts *PossibilititesOpts) init() {

	opts.lowerSquareThreshold, opts.upperSquareThreshold = opts.LowerThreshold*opts.LowerThreshold, opts.UpperThreshold*opts.UpperThreshold
}

// DistanceInvalid returns true when distance is invalid
// ( too big or too small )
//go:nosplit
func (opts *PossibilititesOpts) DistanceInvalid(one, two image.Point) bool {
	distance := geometry.SquareDistance(one, two)
	return distance < opts.lowerSquareThreshold || distance > opts.upperSquareThreshold
}

//trianglesToFirst gives you every possible triangles from points[0]
//go:nosplit
func trianglesToFirst(points geometry.Points, opts PossibilititesOpts) []Triangle {
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
			if area < opts.MinArea {
				continue
			}
			res = append(res, t)
		}
	}
	return res
}

// AllPossibilities returns every possible triangle from the points.
func AllPossibilities(opts PossibilititesOpts, points geometry.Points) []Triangle {
	opts.init()
	if opts.LowerThreshold == opts.UpperThreshold {
		log.Println("EveryTriangles: Identical tresholds, this is not going to work.")
		return nil
	}
	// sort points in order to have cpu
	// do faster using branching
	sort.Sort(points)

	res := []Triangle{}
	for i := 0; i < len(points); i++ {
		tri := trianglesToFirst(points[i:], opts)
		if len(tri) != 0 {
			res = append(res, tri...)
		}
	}
	return res
}
