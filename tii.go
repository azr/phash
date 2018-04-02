package phash

import (
	"image"

	"github.com/azr/phash/cmd"
	"github.com/disintegration/imaging"
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

// GetImageHashesForTriangles calls GetImageHashesForTriangle for each triangle
func GetImageHashesForTriangles(src image.Image, triangles []Triangle) []uint64 {
	res := make([]uint64, 0, len(triangles)*NRotations)

	for _, triangle := range triangles {
		hashes := GetImageHashesForTriangle(src, triangle)
		res = append(res, hashes...)
		break
	}
	return res
}

// GetImageHashesForTriangle :
// 1. get image fragment from image using triangle coordinate
// 2. resize fragment to FragmentWidth/FragmentHeight
// 3. computes a dtc for 3 rotation of that
func GetImageHashesForTriangle(src image.Image, triangle Triangle) []uint64 {
	hashes := make([]uint64, NRotations)

	fragment := triangle.ExtractEquilateralTriangleFrom(src)
	for i := 0; i < NRotations; i++ {
		cmd.WriteImageToPath(fragment, "fragment.of.cat")
		hashes[i] = DTC(fragment)
		break
	}
	return hashes
}

// GetImageFragment returns a crop
// of the coordinates of t from src.
// It will in fact be a rectangle of t
// coordinates.
func (t *Triangle) GetImageFragment(src image.Image) image.Image {
	minX := min(t.A().X, min(t.B().X, t.C().X))
	minY := min(t.A().Y, min(t.B().Y, t.C().Y))
	maxX := max(t.A().X, max(t.B().X, t.C().X))
	maxY := max(t.A().Y, max(t.B().Y, t.C().Y))

	// TODO(azr): do we need to make bottom line of triangle horizontal ?
	// I Don't think so
	return imaging.Crop(src, image.Rect(minX, minY, maxX, maxY))
}
