package geometry

import (
	"image"
	"image/color"
)

type matrix2d [4]int

func (m2d matrix2d) mul(x, y int) (int, int) {
	return m2d[0]*x + m2d[1]*y,
		m2d[2]*x + m2d[3]*y
}

// image2DTransform is a 2 by 2 matrix for 2d image manips
// it allows to transform images while accessing coordinates.
type image2DTransformInt struct {
	m      matrix2d
	src    image.Image
	bounds image.Rectangle
}

// At gives you transformed location of x,y
func (i2dt *image2DTransformInt) At(x, y int) color.Color {
	xx, yy := i2dt.m.mul(x, y)
	return i2dt.src.At(xx, yy)
}

func (i2dt *image2DTransformInt) Bounds() image.Rectangle { return i2dt.bounds }
func (i2dt *image2DTransformInt) ColorModel() color.Model { return i2dt.src.ColorModel() }

// InPlaceRotation90 returns an image that references src
// and that will output you a 90 degrees rotation of src.
func InPlaceRotation90(src image.Image) image.Image {
	min, max := src.Bounds().Min, src.Bounds().Max

	// THOUGHTS(azr)
	// weird/hard, this works but
	// I have a wrong feeling about it
	maxY := 0 //max.X - max.X
	minY := -min.X - max.X
	maxX := max.Y
	minX := min.Y
	rect := image.Rect(minX, minY, maxX, maxY)

	return &image2DTransformInt{
		bounds: rect,
		src:    src,
		m:      rota90M,
	}
}

var (
	rota90M = matrix2d{
		0, -1,
		1, 0,
	}
)
