package manipulator

import (
	"code.google.com/p/biogo.matrix"
	"image"
	// "image/color"
	// "image/draw"
)

func CropMatrix(src matrix.Matrix, x0, y0, x1, y1 int) (*matrix.Dense, error) {
	mtx := make([][]float64, x1-x0+1)
	for x := x0; x <= x1; x++ {
		mtx[x] = make([]float64, y1-y0+1)
		for y := y0; y <= y1; y++ {
			mtx[x][y] = src.At(x, y)
		}
	}
	return matrix.NewDense(mtx)
}

func GrayImageToMatrix(src image.Gray) (*matrix.Dense, error) {
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

func ImageToMatrix(src image.Image) (*matrix.Dense, error) {
	bounds := src.Bounds()
	mtx := make([][]float64, bounds.Max.X)
	for x := 0; x < bounds.Max.X; x++ {
		mtx[x] = make([]float64, bounds.Max.Y)
		for y := 0; y < bounds.Max.Y; y++ {
			mtx[x][y] = ColorToFloat64(src.At(x, y))
		}
	}
	return matrix.NewDense(mtx)
}
