package phash

import (
	"code.google.com/p/biogo.matrix"
	"image"
	"image/color"
	"image/draw"
)

func HammingDistance(hash1, hash2 uint64) uint64 {
	x := hash1 ^ hash2
	var m1, m2, h01, m4 uint64 = 0x5555555555555555, 0x3333333333333333, 0x0101010101010101, 0x0f0f0f0f0f0f0f0f
	x -= (x >> 1) & m1
	x = (x & m2) + ((x >> 2) & m2)
	x = (x + (x >> 4)) & m4
	return (x * h01) >> 56
}

func roundingFactor(x float64) float64 {
	if (x) >= 0 {
		return 0.5
	}
	return -0.5
}

func cropMatrix(src matrix.Matrix, x0, y0, x1, y1 int) (*matrix.Dense, error) {
	mtx := make([][]float64, x1-x0+1)
	for x := x0; x <= x1; x++ {
		mtx[x] = make([]float64, y1-y0+1)
		for y := y0; y <= y1; y++ {
			mtx[x][y] = src.At(x, y)
		}
	}
	return matrix.NewDense(mtx)
}

func grayImageToMatrix(src image.Gray) (*matrix.Dense, error) {
	bounds := src.Bounds()
	mtx := make([][]float64, bounds.Max.X)
	for x := 0; x < bounds.Max.X; x++ {
		mtx[x] = make([]float64, bounds.Max.Y)
		for y := 0; y < bounds.Max.Y; y++ {
			_, _, b, _ := src.At(x, y).RGBA()
			mtx[x][y] = float64(b)
		}
	}
	return matrix.NewDense(mtx)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func Gscl(src image.Image) image.Gray {
	// Create a new grayscale image
	bounds := src.Bounds()
	gray := image.NewGray(bounds)
	for x := 0; x < bounds.Max.X; x++ {
		for y := 0; y < bounds.Max.Y; y++ {
			oldColor := src.At(x, y)
			gray.Set(x, y, color.GrayModel.Convert(oldColor))
		}
	}
	return *gray
}

func CopyImage(src image.Image) draw.Image {
	b := src.Bounds()
	copy := image.NewRGBA(b)

	draw.Draw(copy, copy.Bounds(), src, b.Min, draw.Src)

	return copy
}
