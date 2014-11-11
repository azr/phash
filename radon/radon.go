package radon

import (
	"code.google.com/p/biogo.matrix"
	"code.google.com/p/graphics-go/graphics"
	"errors"
	"github.com/azr/phash/floats"
	"github.com/azr/phash/manipulator"
	"github.com/nfnt/resize"
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

//BackProjectGray computes back projection of img
// in Gray16 by performing an addition
// of backprojection by line.
// 16Gray avoids white noise.
func BackProjectGray(img image.Gray) (*image.Gray16, error) {
	size := img.Bounds().Size()
	width := size.X
	nbProj := size.Y
	step := 180.0 / float64(nbProj)

	out := image.NewGray16(image.Rect(0, 0, width, width))

	for y := 0; y < nbProj; y++ {
		//Extract a 1D-projection (one row Y of sinogram)
		//                             nw[x,y], se[x,y]
		line := img.SubImage(image.Rect(0, y, width, y+1)).(*image.Gray)

		// 3- Do the backprojection and rotate accordingly
		wideLine := resize.Resize(uint(width), uint(width), line, resize.Lanczos3).(*image.Gray)

		θ := manipulator.Rad(float64(-y) * step)
		rotatedWideLine := image.NewGray(image.Rect(0, 0, width, width))
		err := graphics.Rotate(rotatedWideLine, wideLine, &graphics.RotateOptions{Angle: θ})
		if err != nil {
			return out, err
		}

		// 4- Add the rotated backprojection in the output image
		for x := 0; x < width; x++ {
			for y := 0; y < width; y++ {
				point := uint16(out.At(x, y).(color.Gray16).Y) + uint16(rotatedWideLine.At(x, y).(color.Gray).Y)
				out.Set(x, y, color.Gray16{uint16(point)})
			}
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
	Σ := 0.0
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
		Σ += featV[k]
		sumSqd += featV[k] * featV[k]
	}
	mean := Σ / float64(N)
	variable := math.Sqrt((sumSqd / float64(N)) - (Σ*Σ)/float64(N*N))

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
		Σ := 0.0
		for n := 0; n < N; n++ {
			temp := R[n] * math.Cos((math.Pi*float64((2*n+1)*k))/float64(2*N))
			Σ += temp
		}
		if k == 0 {
			DTemp[k] = Σ / math.Sqrt(float64(N))
		} else {
			DTemp[k] = Σ * math.Sqrt(2) / math.Sqrt(float64(N))
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

//AutoCorrelateProjections calculates a cross-correlation of each radial
//projection with itself
func AutoCorrelateProjections(projections image.Gray) []float64 {
	size := projections.Bounds().Size()
	width := size.X
	nbProj := size.Y

	out := make([]float64, nbProj*width)

	// for each projection
	for θ := 0; θ < nbProj; θ++ {
		projection := projections.Pix[projections.PixOffset(0, θ):projections.PixOffset(width, θ)]
		out = append(out, AutoCorrelateSeries(projection)...)
	}

	return out
}

// CrossCorrelate Computes the correlation of two signals at delay lag
//
// param xCoeffs []uint8
// param xCoeffs []uint8
// param lag     int
//
// returns r(the cross correlation array), error(if failed)
func CrossCorrelate(xCoeffs, yCoeffs []uint8, lag int) (float64, error) {

	if len(yCoeffs) != len(xCoeffs) {
		return 0.0, errors.New("signals have different len")
	}
	N := len(yCoeffs)

	mx, my := mean(xCoeffs), mean(yCoeffs)

	Σxy := 0.0
	for i := 0; i < N; i++ {
		j := (i + lag) % N
		Σxy += (float64(xCoeffs[i]) - mx) * (float64(yCoeffs[j]) - my)
	}

	return Σxy / denom(xCoeffs, yCoeffs, mx, my), nil
}

// CrossCorrelateSeries Computes CrossCorrelation with lags
// from 0 to len(Coeffs) of two vectors
// Giving back a correlogram, or an autocorrelation series
//
// param Coeffs []uint8
// param lag int
//
// returns r(the cross correlation)
func CrossCorrelateSeries(x, y []uint8) ([]float64, error) {
	if len(y) != len(x) {
		return nil, errors.New("signals have different len")
	}

	N := len(x)

	correlogram := make([]float64, N)

	for n := 0; n < N; n++ {
		correlogram[n], _ = CrossCorrelate(x, y, n)
	}

	return correlogram, nil
}

// Calculate the mean of a serie x[]
func mean(x []uint8) float64 {
	N := len(x)
	Σ := 0.0
	for i := 0; i < N; i++ {
		Σ += float64(x[i])
	}
	return Σ / float64(N)
}

// Calculate the denominator
func denom(x, y []uint8, mx, my float64) float64 {
	N := len(x)
	Σx := 0.0
	Σy := 0.0
	for i := 0; i < N; i++ {
		Σx += math.Pow(float64(x[i])-mx, 2)
		Σy += math.Pow(float64(y[i])-my, 2)
	}
	return math.Sqrt(Σx * Σy)
}

// AutoCorrelate Computes the cross correlation of a signal
// with itself at lag, circularly.
//
// param Coeffs []uint8
// param lag int
//
// returns r(the cross correlation)
func AutoCorrelate(Coeffs []uint8, lag int) float64 {
	N := len(Coeffs)
	mx := mean(Coeffs)
	Σup, Σpow := 0.0, 0.0

	for n := 0; n < N; n++ {
		h := (n + lag) % N
		x, y := float64(Coeffs[n])-mx, float64(Coeffs[h])-mx
		Σup += (x * y)
		Σpow += math.Pow(x, 2)
	}

	return Σup / Σpow
}

// AutoCorrelateSeries Computes AutoCorrelation with lags
// from 0 to len(Coeffs)
// Giving back a correlogram, or an autocorrelation series
//
// param Coeffs []uint8
// param lag int
//
// returns r(the cross correlation)
func AutoCorrelateSeries(Coeffs []uint8) []float64 {
	N := len(Coeffs)

	correlogram := make([]float64, N)

	for n := 0; n < N; n++ {
		correlogram[n] = AutoCorrelate(Coeffs, n)
	}

	return correlogram
}

// DiffByCrossCorr Computes the cross correlation of two series vectors
// and tells if signals are similar
// param xCoeffs []uint8
// param xCoeffs []uint8
// param threshold - threshold value for which 2 signals are considered the same or different.
//
// returns (true for similar, false for different), (the peak of cross correlation)
func DiffByCrossCorr(xCoeffs, yCoeffs []uint8, threshold float64) (bool, float64) {
	r, err := CrossCorrelateSeries(xCoeffs, yCoeffs)
	max := 0.0
	if err != nil {
		return false, max
	}
	for d := 0; d < len(r); d++ {
		max = math.Max(max, r[d])
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

//ToImage transforms back a set of projections
//in an image
func (p *Projections) ToImage() image.Image {
	return manipulator.MatrixToImage(p.R)
}

//LogMap computes log mapping of signal
// Todo: http://goo.gl/oEiLUh
func LogMap(in []float64) []float64 {
	Σ := floats.Sum(in)
	out := make([]float64, len(in))

	for x := 0; x < len(in); x++ {
		out[x] = math.Log(Σ * math.Exp(in[x]))
	}
	return out
}
