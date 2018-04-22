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
// If two perceptual hashes are equal it means that the features
// in the triangle are perceptually similar.
//
// Triangles could come from FindKeypoints or your own library.
func FromTriangles(src image.Image, triangles []triangle.Triangle) <-chan uint64 {
	c := make(chan uint64)

	go func() {
		parallel(len(triangles), func(start, end int) {
			for i := start; i < end; i++ {
				triangle := triangles[i]
				fragment := triangle.ExtractEquilateralTriangleFrom(src)
				for i := 0; i < rotations; i++ {
					c <- DTC(fragment)
					fragment = geometry.InPlaceRotation90(fragment)
				}
			}
		})
		close(c)
	}()
	return c
}
