package phash

import (
	"code.google.com/p/biogo.matrix"
	"image"
	"math"
)

const (
	nb_coeffs_radon = 80
)

type Projections struct {
	R              matrix.Matrix //contains projections of image of angled lines through center
	nb_pix_perline []int         //the head of int array denoting the number of pixels of each line
	size           int           //the size of nb_pix_perline
}

type Features struct {
	features []float64 //the head of the feature array of double's
	size     int       //the size of the feature array
}

type Digest struct {
	coeffs []uint8 //the head of the digest integer coefficient array
	size   int     //the size of the coeff array
}

func radonProjections(img matrix.Matrix, N int) (Projections, error) {
	var projs Projections
	var err error
	width, height := img.Dims()

	D := max(width, height)
	var x_center, y_center float64 = float64(width) / 2.0, float64(height) / 2.0

	x_off := int(math.Floor(x_center + roundingFactor(x_center)))
	y_off := int(math.Floor(y_center + roundingFactor(y_center)))

	radon_map, err := matrix.ZeroDense(N, D)
	if err != nil {
		return projs, err
	}
	projs.R = radon_map

	projs.nb_pix_perline = make([]int, N)
	projs.size = N
	nb_per_line := projs.nb_pix_perline

	for k := 0; k < N/4+1; k++ {
		theta := float64(k) * math.Pi / float64(N)
		alpha := math.Tan(theta)
		for x := 0; x < D; x++ {

			y := alpha * float64(x-x_off)
			yd := int(math.Floor(y + roundingFactor(y)))

			if (yd+y_off >= 0) && (yd+y_off < height) && (x < width) {
				radon_map.Set(k, x, img.At(x, yd+y_off))
				nb_per_line[k] += 1
			}
			if (yd+x_off >= 0) && (yd+x_off < width) && (k != N/4) && (x < height) {
				radon_map.Set(N/2-k, x, img.At(yd+x_off, x))
				nb_per_line[N/2-k] += 1
			}

		}
	}
	j := 0
	for k := 3 * N / 4; k < N; k++ {
		theta := float64(k) * math.Pi / float64(N)
		alpha := math.Tan(theta)
		for x := 0; x < D; x++ {
			y := alpha * float64(x-x_off)
			yd := int(math.Floor(y + roundingFactor(y)))
			if (yd+y_off >= 0) && (yd+y_off < height) && (x < width) {
				radon_map.Set(k, x, img.At(x, yd+y_off))
				nb_per_line[k] += 1
			}
			if (y_off-yd >= 0) && (y_off-yd < width) && (2*y_off-x >= 0) && (2*y_off-x < height) && (k != 3*N/4) {
				radon_map.Set(k-j, x, img.At(-yd+y_off, -(x-y_off)+y_off))
				nb_per_line[k-j] += 1
			}

		}
		j += 2
	}

	return projs, nil

}

func featureVector(projs Projections) Features {
	var fv Features
	projection_map := projs.R
	nb_perline := projs.nb_pix_perline
	N := projs.size
	_, D := projection_map.Dims()

	fv.features = make([]float64, N)
	fv.size = N

	feat_v := fv.features
	sum := 0.0
	sum_sqd := 0.0
	for k := 0; k < N; k++ {
		line_sum := 0.0
		line_sum_sqd := 0.0
		nb_pixels := nb_perline[k]
		for i := 0; i < D; i++ {
			line_sum += projection_map.At(k, i)
			line_sum_sqd += projection_map.At(k, i) * projection_map.At(k, i)
		}
		feat_v[k] = (line_sum_sqd / float64(nb_pixels)) - (line_sum*line_sum)/float64(nb_pixels*nb_pixels)
		sum += feat_v[k]
		sum_sqd += feat_v[k] * feat_v[k]
	}
	mean := sum / float64(N)
	variable := math.Sqrt((sum_sqd / float64(N)) - (sum*sum)/float64(N*N))

	for i := 0; i < N; i++ {
		feat_v[i] = (feat_v[i] - mean) / variable
	}

	return fv
}

func dct(fv Features) Digest {
	var digest Digest

	N := fv.size
	nb_coeffs := nb_coeffs_radon

	digest.coeffs = make([]uint8, nb_coeffs)

	digest.size = nb_coeffs

	R := fv.features

	D := digest.coeffs

	var D_temp [nb_coeffs_radon]float64
	max := 0.0
	min := 0.0

	for k := 0; k < nb_coeffs; k++ {
		sum := 0.0
		for n := 0; n < N; n++ {
			temp := R[n] * math.Cos((math.Pi*float64((2*n+1)*k))/float64(2*N))
			sum += temp
		}
		if k == 0 {
			D_temp[k] = sum / math.Sqrt(float64(N))
		} else {
			D_temp[k] = sum * math.Sqrt(2) / math.Sqrt(float64(N))
		}
		if D_temp[k] > max {
			max = D_temp[k]
		}
		if D_temp[k] < min {
			min = D_temp[k]
		}

	}

	for i := 0; i < nb_coeffs; i++ {
		D[i] = uint8(math.MaxUint8 * (D_temp[i] - min) / (max - min))
	}

	return digest
}

func crosscorr(x, y Digest, threshold float64) bool {

	N := y.size

	x_coeffs := x.coeffs
	y_coeffs := y.coeffs

	r := make([]float64, N)
	sumx := 0.0
	sumy := 0.0
	for i := 0; i < N; i++ {
		sumx += float64(x_coeffs[i])
		sumy += float64(y_coeffs[i])
	}
	meanx := sumx / float64(N)
	meany := sumy / float64(N)
	max := 0.0
	for d := 0; d < N; d++ {
		num := 0.0
		denx := 0.0
		deny := 0.0
		for i := 0; i < N; i++ {
			num += (float64(x_coeffs[i]) - meanx) * (float64(y_coeffs[(N+i-d)%N]) - meany)
			denx += math.Pow((float64(x_coeffs[i]) - meanx), 2)
			deny += math.Pow((float64(y_coeffs[(N+i-d)%N]) - meany), 2)
		}
		r[d] = num / math.Sqrt(denx*deny)
		if r[d] > float64(max) {
			max = r[d]
		}
	}

	return max > threshold
}

func Radon(img image.Gray) Digest {
	imgMtx, err := getImageMatrix(img)
	if err != nil {
		panic(err)
	}
	radonProjection, err := radonProjections(imgMtx, img.Bounds().Max.X)
	if err != nil {
		panic(err)
	}

	fv := featureVector(radonProjection)
	dctDigest := dct(fv)
	return dctDigest
}
