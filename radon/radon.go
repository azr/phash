package radon

import (
	"code.google.com/p/graphics-go/graphics"
	"errors"
	"github.com/azr/phash/floats"
	"github.com/azr/phash/manipulator"
	"github.com/disintegration/imaging"
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
	R            *image.Gray //projections
	nbPixPerline []uint8     //the head of int array denoting the number of pixels of each line
}

//FeaturesVector contains feature vector info
type FeaturesVector struct {
	features []float64 //the head of the feature array of double's
}

type Digest struct {
	Coeffs []uint8 //the head of the Radondigest integer coefficient array
}

func addPixelsToGray(src image.Image, sx, sy int, dst image.Gray, dx, dy int) {
	clr := src.At(sx, sy)
	greyColor, _ := color.GrayModel.Convert(clr).(color.Gray)
	dst.Set(dx, dy, color.Gray{dst.At(dx, dy).(color.Gray).Y + greyColor.Y})
}

// Project finds radon projections of N lines running through the image
// center for lines angled 0 to 180 degrees from horizontal.
// Totally not working
//
//  /param img - img.Image src image
//  /param  N  - int number of angled lines to consider.
//  /return sinogram - Projections
//  /return nbPerLine - number of pixel per line in sinogram
//  /return error - if failed
func Project(img image.Image, N int) (*image.Gray, []uint8, error) {
	if N == 0 {
		N = 180
	}
	var err error
	size := img.Bounds().Size()

	D := max(size.X, size.Y)
	xCenter, yCenter := float64(size.X)/2.0, float64(size.Y)/2.0

	xOff := int(math.Floor(xCenter + roundingFactor(xCenter)))
	yOff := int(math.Floor(yCenter + roundingFactor(yCenter)))

	radonMap := image.NewGray(image.Rect(0, 0, N, D))
	if err != nil {
		return nil, nil, err
	}

	nbPerLine := make([]uint8, N)

	for k := 0; k < (N/4)+1; k++ {
		θ := float64(k) * math.Pi / float64(N)
		α := math.Tan(θ)
		for x := 0; x < D; x++ {

			y := α * float64(x-xOff)
			yd := int(math.Floor(y + roundingFactor(y)))

			if (yd+yOff >= 0) && (yd+yOff < size.Y) && (x < size.X) {
				addPixelsToGray(img, x, yd+yOff, *radonMap, k, x)
				nbPerLine[k]++
			}
			if (yd+xOff >= 0) && (yd+xOff < size.X) && (k != N/4) && (x < size.Y) {
				addPixelsToGray(img, yd+xOff, x, *radonMap, N/2-k, x)
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
			if (yd+yOff >= 0) && (yd+yOff < size.Y) && (x < size.X) {
				addPixelsToGray(img, x, yd+yOff, *radonMap, k, x)
				nbPerLine[k]++
			}
			if (yOff-yd >= 0) && (yOff-yd < size.X) && (2*yOff-x >= 0) && (2*yOff-x < size.Y) && (k != 3*N/4) {
				addPixelsToGray(img, -yd+yOff, -(x-yOff)+yOff, *radonMap, k-j, x)
				nbPerLine[k-j]++
			}

		}
		j += 2
	}

	return radonMap, nbPerLine, nil
}

func atAngle(img image.Image, θ float64, x, y int, xc, yc float64) color.Color {
	xθ := floats.Round(math.Cos(θ)*(float64(x)-xc)-math.Sin(θ)*(float64(y)-yc)+xc, 0.5, 0)
	yθ := floats.Round(math.Cos(θ)*(float64(y)-yc)-math.Sin(θ)*(float64(x)-xc)+yc, 0.5, 0)
	// log.Println("yah :", xθ, yθ)
	return img.At(int(xθ+roundingFactor(xθ)), int(yθ+roundingFactor(yθ)))
}

func getSinogramFor(θ float64, img image.Image) []float64 {
	sinogram := make([]float64, img.Bounds().Size().X)
	xc := float64(img.Bounds().Size().X) / 2
	yc := float64(img.Bounds().Size().Y) / 2

	// get column average profile
	for y := 0; y < img.Bounds().Size().Y; y++ {
		for x := 0; x < img.Bounds().Size().X; x++ {
			clr := atAngle(img, θ, x, y, xc, yc)
			greyColor, _ := color.GrayModel.Convert(clr).(color.Gray)
			sinogram[x] = sinogram[x] + float64(greyColor.Y)
		}
	}
	return sinogram
}

//ProjectGray returns a greyscale sinogram (or radon projection) of img
// N(default: 360): number of image rotation on which a projection will be done
// A naive simplistic approach
// Sinograms looks like this :
// θ1 θ2 θ3...θN
// |  |  |    |
// |  |  |    |
func ProjectGray(src image.Image, N int) (*image.Gray, error) {
	if N == 0 {
		N = 360
	}
	step := 180.0 / float64(N)

	size := src.Bounds().Size()
	overX := int(float64(size.X) * 1.1)
	overY := int(float64(size.Y) * 1.1)
	var img image.Image = image.NewGray(image.Rect(0, 0, size.X+overX, size.Y+overY))
	img = imaging.Overlay(img, src, image.Pt(overX/2, overY/2), 1)
	size = img.Bounds().Size()

	D := max(size.X, size.Y)
	out := image.NewGray(image.Rect(0, 0, N, D))

	// for each given angle θ
	for n := 0; n < N; n++ {
		θ := float64(n) * step

		sinogram := getSinogramFor(θ, img)

		//Set out line with sinogram
		for d := 0; d < D; d++ {
			out.Set(n, d, color.Gray{uint8(sinogram[d] / float64(size.Y))})
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
	width := size.Y
	nbProj := size.X
	step := 180.0 / float64(nbProj)

	out := image.NewGray16(image.Rect(0, 0, width, width))

	for X := 0; X < nbProj; X++ {
		//Extract a 1D-projection (one row Y of sinogram)
		line := img.SubImage(image.Rect(X, 0, X+1, width)).(*image.Gray)

		// 3- Do the backprojection and rotate accordingly
		wideLine := resize.Resize(uint(width), uint(width), line, resize.Lanczos3).(*image.Gray)

		θ := manipulator.Rad(float64(X)*step) + math.Pi/2
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
	D := projectionMap.Bounds().Size().X

	fv.features = make([]float64, N)

	featV := fv.features
	Σ := 0.0
	sumSqd := 0.0
	for k := 0; k < N; k++ {
		var lineSum, lineSumSqd uint8
		nbPixels := nbPerline[k]
		for i := 0; i < D; i++ {
			lineSum += projectionMap.At(k, i).(color.Gray).Y
			lineSumSqd += projectionMap.At(k, i).(color.Gray).Y * projectionMap.At(k, i).(color.Gray).Y
		}
		featV[k] = float64(lineSumSqd)/float64(nbPixels) - float64(lineSum*lineSum)/float64(nbPixels*nbPixels)
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
		// log.Println("θ:", θ)
		left := projections.PixOffset(0, θ)
		right := projections.PixOffset(width, θ)
		projection := projections.Pix[left:right]
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
func DigestMatrix(img image.Image) (Digest, Projections, FeaturesVector) {
	var radonProjection Projections
	var err error
	radonProjection.R, radonProjection.nbPixPerline, err = Project(img, 0)
	if err != nil {
		panic(err)
	}

	fv := FeatureVector(radonProjection)
	dctDigest := Dct(fv)
	return dctDigest, radonProjection, fv
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
