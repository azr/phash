package manipulator

import (
	"code.google.com/p/biogo.matrix"
	"image"
	"image/color"
	"image/draw"
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

//ImageToGrayscale returns the greyscale of an image
func ImageToGrayscale(src image.Image) image.Gray {
	// Create a new grayscale image
	bounds := src.Bounds()
	gray := image.NewGray(bounds)
	for x := 0; x < bounds.Max.X; x++ {
		for y := 0; y < bounds.Max.Y; y++ {
			gray.Set(x, y, color.GrayModel.Convert(src.At(x, y)))
		}
	}
	return *gray
}

func MatrixToImage(src matrix.Matrix) image.Image {
	width, height := src.Dims()
	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			img.Set(x, y, Float64ToColor(src.At(x, y)))
		}
	}
	return img
}

//CopyImage copies images into a draw.Image
func CopyImage(src image.Image) draw.Image {
	b := src.Bounds()
	copy := image.NewRGBA(b)

	draw.Draw(copy, copy.Bounds(), src, b.Min, draw.Src)

	return copy
}
