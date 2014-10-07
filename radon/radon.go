package radon

import (
	"code.google.com/p/biogo.matrix"
	"code.google.com/p/graphics-go/graphics"
	"github.com/azr/phash/manipulator"
	"image"
	"image/color"
	"math"
)

const (
	nbCoeffsRadon = 80
)

//Projections contains radon projections of image of angled lines through center
type Projections struct {
	R            matrix.Matrix //projections
	nbPixPerline []int         //the head of int array denoting the number of pixels of each line
}

//FeaturesVector contains feature vector info
type FeaturesVector struct {
	features []float64 //the head of the feature array of double's
}

type Digest struct {
	Coeffs []uint8 //the head of the Radondigest integer coefficient array
}

// Project finds radon projections of N lines running through the image
// center for lines angled 0to 180 degrees from horizontal.
//  /param img - img [Matrix](code.google.com/p/biogo.matrix) src image
//  /param  N  - int number of angled lines to consider.
//  /return projs - Projections struct
//  /return error - if failed
func Project(img matrix.Matrix, N int) (Projections, error) {
	var projs Projections
	var err error
	width, height := img.Dims()

	if N == 0 {
		N = 180
	}
	D := max(width, height)
	var xCenter, yCenter float64 = float64(width) / 2.0, float64(height) / 2.0

	xOff := int(math.Floor(xCenter + roundingFactor(xCenter)))
	yOff := int(math.Floor(yCenter + roundingFactor(yCenter)))

	radonMap, err := matrix.ZeroDense(N, D)
	if err != nil {
		return projs, err
	}
	projs.R = radonMap

	projs.nbPixPerline = make([]int, N)
	nbPerLine := projs.nbPixPerline

	for k := 0; k < N/4+1; k++ {
		theta := float64(k) * math.Pi / float64(N)
		alpha := math.Tan(theta)
		for x := 0; x < D; x++ {

			y := alpha * float64(x-xOff)
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
	j := 0
	for k := 3 * N / 4; k < N; k++ {
		theta := float64(k) * math.Pi / float64(N)
		alpha := math.Tan(theta)
		for x := 0; x < D; x++ {
			y := alpha * float64(x-xOff)
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

	return projs, nil
}

//ProjectGray returns a greyscale sinogram (or radon projection) of img
// N(default: 180): number of image rotation on which a projection will be done
func ProjectGray(img image.Image, N int) (*image.Gray, error) {
	if N == 0 {
		N = 180
	}
	step := 180.0 / float64(N)

	size := img.Bounds().Size()
	D := max(size.X, size.Y)
	out := image.NewGray(image.Rect(0, 0, D, N))

	// for each given angle θ
	for y := 0; y < N; y++ {
		θ := float64(y) * step
		draw := image.NewRGBA(image.Rect(0, 0, img.Bounds().Dy(), img.Bounds().Dx()))
		//have a duplicate img rotated by θ
		err := graphics.Rotate(draw, img, &graphics.RotateOptions{Angle: manipulator.Rad(θ)})
		if err != nil {
			return out, err
		}

		// compute projection of image
		sinogram := make([]float64, draw.Bounds().Size().X)
		// get column average profile
		for y := 0; y < draw.Bounds().Size().Y; y++ {
			for x := 0; x < draw.Bounds().Size().X; x++ {
				greyColor, _ := color.GrayModel.Convert(draw.At(x, y)).(color.Gray)
				sinogram[x] = sinogram[x] + float64(greyColor.Y)
			}
		}

		//Set out line with sinogram
		for x := 0; x < out.Bounds().Size().X; x++ {
			out.Set(x, y, color.Gray{uint8(sinogram[x] / float64(draw.Bounds().Size().X))})
		}
	}

	return out, nil
}

//FeatureVector computes the feature vector from a radon projection map.
func FeatureVector(projs Projections) FeaturesVector {
	var fv FeaturesVector
	projectionMap := projs.R
	nbPerline := projs.nbPixPerline
	N := len(projs.nbPixPerline)
	_, D := projectionMap.Dims()

	fv.features = make([]float64, N)

	featV := fv.features
	sum := 0.0
	sumSqd := 0.0
	for k := 0; k < N; k++ {
		lineSum := 0.0
		lineSumSqd := 0.0
		nbPixels := nbPerline[k]
		for i := 0; i < D; i++ {
			lineSum += projectionMap.At(k, i)
			lineSumSqd += projectionMap.At(k, i) * projectionMap.At(k, i)
		}
		featV[k] = (lineSumSqd / float64(nbPixels)) - (lineSum*lineSum)/float64(nbPixels*nbPixels)
		sum += featV[k]
		sumSqd += featV[k] * featV[k]
	}
	mean := sum / float64(N)
	variable := math.Sqrt((sumSqd / float64(N)) - (sum*sum)/float64(N*N))

	for i := 0; i < N; i++ {
		featV[i] = (featV[i] - mean) / variable
	}

	return fv
}

//Dct Computes the radon digest of a given vector
func Dct(fv FeaturesVector) Digest {
	var Digest Digest

	N := len(fv.features)
	nbCoeffs := nbCoeffsRadon

	Digest.Coeffs = make([]uint8, nbCoeffs)

	R := fv.features

	D := Digest.Coeffs

	var DTemp [nbCoeffsRadon]float64
	max := 0.0
	min := 0.0

	for k := 0; k < nbCoeffs; k++ {
		sum := 0.0
		for n := 0; n < N; n++ {
			temp := R[n] * math.Cos((math.Pi*float64((2*n+1)*k))/float64(2*N))
			sum += temp
		}
		if k == 0 {
			DTemp[k] = sum / math.Sqrt(float64(N))
		} else {
			DTemp[k] = sum * math.Sqrt(2) / math.Sqrt(float64(N))
		}
		if DTemp[k] > max {
			max = DTemp[k]
		}
		if DTemp[k] < min {
			min = DTemp[k]
		}

	}

	for i := 0; i < nbCoeffs; i++ {
		D[i] = uint8(math.MaxUint8 * (DTemp[i] - min) / (max - min))
	}

	return Digest
}

// CrossCorr Computes the cross correlation of two series vectors
// param x - Digest struct
// param y - Digest struct
// param threshold - threshold value for which 2 images are considered the same or different.
//
// returns (true for similar, false for different), (the peak of cross correlation)
func CrossCorr(x, y Digest, threshold float64) (bool, float64) {

	N := len(y.Coeffs)

	xCoeffs := x.Coeffs
	yCoeffs := y.Coeffs

	r := make([]float64, N)
	sumx := 0.0
	sumy := 0.0
	for i := 0; i < N; i++ {
		sumx += float64(xCoeffs[i])
		sumy += float64(yCoeffs[i])
	}
	meanx := sumx / float64(N)
	meany := sumy / float64(N)
	max := 0.0
	for d := 0; d < N; d++ {
		num := 0.0
		denx := 0.0
		deny := 0.0
		for i := 0; i < N; i++ {
			num += (float64(xCoeffs[i]) - meanx) * (float64(yCoeffs[(N+i-d)%N]) - meany)
			denx += math.Pow((float64(xCoeffs[i]) - meanx), 2)
			deny += math.Pow((float64(yCoeffs[(N+i-d)%N]) - meany), 2)
		}
		r[d] = num / math.Sqrt(denx*deny)
		if r[d] > float64(max) {
			max = r[d]
		}
	}

	return max > threshold, max
}

//DigestMatrix computes radon digest of img matrix
func DigestMatrix(img *matrix.Dense) (Digest, Projections, FeaturesVector) {
	radonProjection, err := Project(img, 0)
	if err != nil {
		panic(err)
	}

	fv := FeatureVector(radonProjection)
	dctDigest := Dct(fv)
	return dctDigest, radonProjection, fv
}

func (p *Projections) ToImage() image.Image {
	return manipulator.MatrixToImage(p.R)
}
