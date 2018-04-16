package phash

import (
	"image"

	"github.com/azr/phash/cmd"
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
// Triangles could come from FindKeypoints or of your own library.
func FromTriangles(src image.Image, triangles []triangle.Triangle) []uint64 {
	res := make([]uint64, 0, len(triangles)*rotations)

	for _, triangle := range triangles {
		hashes := getImageHashesForTriangle(src, triangle)
		res = append(res, hashes...)
		break
	}
	return res
}

// getImageHashesForTriangle :
//  1. get image fragment from image using triangle coordinate
//  2. resize fragment to FragmentWidth/FragmentHeight
//  3. computes a dtc for 3 rotation of that
func getImageHashesForTriangle(src image.Image, triangle triangle.Triangle) []uint64 {
	hashes := make([]uint64, rotations)

	fragment := triangle.ExtractEquilateralTriangleFrom(src)
	for i := 0; i < rotations; i++ {
		cmd.WriteImageToPath(fragment, "fragment.of.cat")
		hashes[i] = DTC(fragment)
		break
	}
	return hashes
}
