package phash

import (
	"image"
	"math"
)

//Projections contains radon projections of image of angled lines through center
type Projections struct {
	R            image.Image //projections
	nbPixPerline []uint8     //the head of int array denoting the number of pixels of each line in R
}

// RadonProjections computes radon projections of N lines running through the image
// center for lines angled 0 to 180 degrees from horizontal.
//
//  /param img - img.Image src image
//  /param  N  - int number of angled lines to consider.
//  /return sinogram - Projections
//  /return nbPerLine - number of pixel per line in sinogram
func RadonProjections(img image.Image, N int) *Projections {
	if N == 0 {
		N = 180
	}
	size := img.Bounds().Size()
	height, width := size.X, size.Y

	D := max(width, size.Y)
	xCenter, yCenter := float64(width)/2.0, float64(height)/2.0

	xOff := int(math.Floor(xCenter + roundingFactor(xCenter)))
	yOff := int(math.Floor(yCenter + roundingFactor(yCenter)))

	radonMap := image.NewRGBA(image.Rect(0, 0, N, D)) // TODO(azr): create image of the same type as img ?

	nbPerLine := make([]uint8, N)

	for k := 0; k < (N/4)+1; k++ {
		θ := float64(k) * math.Pi / float64(N)
		α := math.Tan(θ)
		for x := 0; x < D; x++ {

			y := α * float64(x-xOff)
			yd := int(math.Floor(y + roundingFactor(y)))

			if (yd+yOff >= 0) && (yd+yOff < height) && (x < width) {
				radonMap.Set(k, x, img.At(x, yd+yOff))
				nbPerLine[k]++
			}
			if (yd+xOff >= 0) && (yd+xOff < width) && (k != N/4) && (x < height) {
				radonMap.Set(N/2-k, x, img.At(yd+xOff, x))
				nbPerLine[N/2-k]++
			}

		}
	}
	for j, k := 0, 3*N/4; k < N; k++ {
		θ := float64(k) * math.Pi / float64(N)
		α := math.Tan(θ)
		for x := 0; x < D; x++ {
			y := α * float64(x-xOff)
			yd := int(math.Floor(y + roundingFactor(y)))
			if (yd+yOff >= 0) && (yd+yOff < height) && (x < width) {
				radonMap.Set(k, x, img.At(x, yd+yOff))
				nbPerLine[k]++
			}
			if (yOff-yd >= 0) && (yOff-yd < width) && (2*yOff-x >= 0) && (2*yOff-x < height) && (k != 3*N/4) {
				radonMap.Set(k-j, x, img.At(-yd+yOff, -(x-yOff)+yOff))
				nbPerLine[k-j]++
			}

		}
		j += 2
	}

	return &Projections{radonMap, nbPerLine}

}
