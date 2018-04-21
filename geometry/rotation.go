package geometry

import (
	"image"
	"image/color"
)

func (i2dt *image2DTransformInt) mul(x, y int) (int, int) {
	return i2dt.m[0]*x + i2dt.m[1]*y,
		i2dt.m[2]*x + i2dt.m[3]*y
}

// image2DTransform is a 2 by 2 matrix for 2d image manips
// it allows to transform images while accessing coordinates.
type image2DTransformInt struct {
	m      [4]int
	src    image.Image
	bounds image.Rectangle
}

// At gives you transformed location of x,y
func (i2dt *image2DTransformInt) At(x, y int) color.Color {
	return i2dt.src.At(i2dt.mul(x, y))
}

func (i2dt *image2DTransformInt) Bounds() image.Rectangle { return i2dt.bounds }
func (i2dt *image2DTransformInt) ColorModel() color.Model { return i2dt.src.ColorModel() }

// InPlaceRotation90 returns an image that references src
// and that will output you a 90 degrees rotation of src.
func InPlaceRotation90(src image.Image) image.Image {
	bds := src.Bounds()

	return &image2DTransformInt{
		bounds: image.Rectangle{
			Min: image.Point{X: 0, Y: 0},
			Max: image.Point{X: bds.Dy(), Y: bds.Dx()},
		},
		src: src,
		m:   rota90M,
	}
}

var (
	rota90M = [4]int{
		0, 1,
		1, 0,
	}
)
