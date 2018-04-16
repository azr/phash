package triangle

import (
	"image"
	"log"
	"sort"

	"github.com/azr/phash/geometry"
)

// PossibilititesOpts is an option for EveryTriangles func
type PossibilititesOpts struct {
	Src image.Image // source image

	// min distance ratio between two points
	// 1 for LowerThreshold would mean a triangle has to
	// contain as much pixel as the image has to be valid
	LowerThresholdRatio, UpperThresholdRatio   float32
	lowerThreshold, upperThreshold             int
	lowerSquareThreshold, upperSquareThreshold int

	MinAreaRatio float32 // ratio in pixels an triangle has to be for it to be valid
	minArea      int     // in pixel ratio, 1 means everything
}

func (opts *PossibilititesOpts) init() {
	bounds := opts.Src.Bounds()
	pixels := (bounds.Max.X - bounds.Min.X) * (bounds.Max.Y - bounds.Min.Y)
	opts.lowerThreshold = int(opts.LowerThresholdRatio * float32(pixels))
	opts.upperThreshold = int(opts.UpperThresholdRatio * float32(pixels))
	opts.minArea = int(opts.MinAreaRatio * float32(pixels))
	opts.lowerSquareThreshold, opts.upperSquareThreshold = opts.lowerThreshold*opts.lowerThreshold, opts.upperThreshold*opts.upperThreshold
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
			if area < opts.minArea {
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
	if opts.lowerThreshold == opts.upperThreshold {
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
