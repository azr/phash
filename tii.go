package phash

import (
	"image"

	"github.com/azr/phash/geometry"
	"github.com/azr/phash/geometry/triangle"
)

const (
	// how many times to rotate a triangle
	// before
	rotations = 3
)

// FromTriangles calculates 3 perceptual hash of src per triangles.
//
// A triangle is transformed into it's equilateral version
// then we run a DTC on the 3 different angles of the triangle.
//
// When two triangle have the same perceptual hash, it means
// that the features in the triangle are perceptually similar.
//
// Triangles could come from FindKeypoints or your own library.
//
// This function will start a goroutine per CPU.
func FromTriangles(src image.Image, triangles []triangle.Triangle) []uint64 {
	res := make([]uint64, len(triangles)*3)

	parallel(len(triangles), func(start, end int) {
		for i := start; i < end; i++ {
			triangle := triangles[i]
			fragment := triangle.ExtractEquilateralTriangleFrom(src)
			for j := 0; j < rotations; j++ {
				res[(i*3)+j] = DTC(fragment)
				fragment = geometry.InPlaceRotation90(fragment)
			}
		}
	})
	return res
}
