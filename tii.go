package phash

import (
	"image"

	"github.com/azr/phash/cmd"
	"github.com/azr/phash/geometry/triangle"
)

const (
	// FragmentWidth is width of a fragment
	// before it's being rotated & hashed
	FragmentWidth = 51
	// FragmentHeight is height of a fragment
	// before it's being rotated & hashed
	FragmentHeight = 60
	// NRotations of a fragment to be hashed
	NRotations = 3
)

// FromTriangles calls GetImageHashesForTriangle for each triangle
func FromTriangles(src image.Image, triangles []triangle.Triangle) []uint64 {
	res := make([]uint64, 0, len(triangles)*NRotations)

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
	hashes := make([]uint64, NRotations)

	fragment := triangle.ExtractEquilateralTriangleFrom(src)
	for i := 0; i < NRotations; i++ {
		cmd.WriteImageToPath(fragment, "fragment.of.cat")
		hashes[i] = DTC(fragment)
		break
	}
	return hashes
}
