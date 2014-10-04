package phash

import (
	"image"
	// "github.com/hawx/img/greyscale"
	"code.google.com/p/biogo.matrix"

	// Package image/[jpeg|fig|png] is not used explicitly in the code below,
	// but is imported for its initialization side-effect, which allows
	// image.Decode to understand [jpeg|gif|png] formatted images.
	_ "code.google.com/p/go.image/bmp"
	_ "code.google.com/p/go.image/tiff"
	_ "code.google.com/p/graphics-go/graphics"
	"github.com/azr/phash/manipulator"
	"github.com/azr/phash/radon"
	// _ "github.com/kavu/go-phash"
	"github.com/nfnt/resize"
	_ "github.com/smartystreets/goconvey/convey"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"math"
	"sort"
)

func median(a []float64) float64 {
	b := make([]float64, len(a))
	copy(b, a)
	sort.Float64s(b)
	return b[int(len(b)/2)]
}

func coefficient(n int) float64 {
	if n == 0 {
		return 1.0 / math.Sqrt(2)
	}
	return 1.0
}

func dctPoint(img image.Gray, u, v, N, M int) float64 {
	sum := 0.0
	for i := 0; i < N; i++ {
		for j := 0; j < M; j++ {
			_, _, b, _ := img.At(i, j).RGBA()
			// sum += math.Cos( ( float64(2*i+1)/float64(2*N) ) * float64( u ) * math.Pi ) *
			//        math.Cos( ( float64(2*j+1)/float64(2*M) ) * float64( v ) * math.Pi ) *
			//        float64(b)

			sum += math.Cos(math.Pi/(float64(N))*(float64(i)+1/2)*float64(u)) *
				math.Cos(math.Pi/(float64(M))*(float64(j)+1/2)*float64(v)) *
				float64(b)
		}
	}
	return sum * ((coefficient(u) * coefficient(v)) / 4.0)
}

// GreyscaleDct Computes the Dct of a greyscale image
func GreyscaleDct(img image.Gray) uint64 {
	// func DctImageHashOne(img image.Image) ([][]float64) {
	R := img.Bounds()
	N := R.Dx() // width
	M := R.Dy() // height
	DCTMatrix := make([][]float64, N)
	for u := 0; u < N; u++ {
		DCTMatrix[u] = make([]float64, M)
		for v := 0; v < M; v++ {
			DCTMatrix[u][v] = dctPoint(img, u, v, N, M)
			// fmt.Println( "DCTMatrix[", u, "][", v, "] is ", DCTMatrix[u][v])
		}
	}

	total := 0.0
	for u := 0; u < N/2; u++ {
		for v := 0; v < M/2; v++ {
			total += DCTMatrix[u][v]
		}
	}
	total -= DCTMatrix[0][0]
	avg := total / float64(((N/2)*(M/2))-1)
	var hash uint64
	for u := 0; u < N/2; u++ {
		for v := 0; v < M/2; v++ {
			hash = hash * 2
			if DCTMatrix[u][v] > avg {
				hash++
			}
		}
	}

	return hash
}

func dctMatrixRow(N, M, x int, c0, c1 float64) []float64 {
	row := make([]float64, M)
	row[0] = c0
	for y := 1; y < M; y++ {
		row[y] = c1 * math.Cos((math.Pi/2.0/float64(N))*float64(y)*(2.0*float64(x)+1.0))
	}

	return row
}

func createDctMatrix(N, M int) (*matrix.Dense, error) {
	mtx := make([][]float64, N)
	c1 := math.Sqrt(2.0 / float64(N))
	c0 := 1 / math.Sqrt(float64(M))
	for x := 0; x < N; x++ {
		mtx[x] = dctMatrixRow(N, M, x, c0, c1)
	}

	return matrix.NewDense(mtx)
}

// GreyscaleDctMatrix Computes the Dct of a greyscale image using matrixes
func GreyscaleDctMatrix(img image.Gray) uint64 {
	imgMtx, err := manipulator.GrayImageToMatrix(img)
	if err != nil {
		panic(err)
	}
	dctMtx, err := createDctMatrix(img.Bounds().Max.X, img.Bounds().Max.Y)
	if err != nil {
		panic(err)
	}
	dctMtxTransp := dctMtx.T(nil) // Transpose

	dctImage := dctMtx.Dot(imgMtx, nil).Dot(dctMtxTransp, nil)

	dctImage, err = manipulator.CropMatrix(dctImage, 0, 0, 7, 7)
	if err != nil {
		panic(err)
	}
	subsec := matrix.ElementsVector(dctImage)
	median := median(subsec)
	var one, hash uint64 = 1, 0
	for i := 0; i < len(subsec); i++ {
		current := subsec[i]
		if current > median {
			hash |= one
		}
		one = one << 1
	}
	return hash
}

//ComputeGreyscaleDct puts the result of GreyscaleDct in a digest
func (d *ImageDigest) ComputeGreyscaleDct() error {
	stamp := resize.Resize(32, 32, d.Radon.Image, resize.Bilinear)
	greyscaleStamp := Gscl(stamp)

	// greyscaleStamp := greyscale.Greyscale(stamp)
	d.Phash = GreyscaleDct(greyscaleStamp)

	return nil
}

//ComputeGreyscaleDctMatrix puts the result of GreyscaleDctMatrix in a digest
func (d *ImageDigest) ComputeGreyscaleDctMatrix() error {
	stamp := resize.Resize(32, 32, d.Radon.Image, resize.Bilinear)
	greyscaleStamp := Gscl(stamp)

	// greyscaleStamp := greyscale.Greyscale(stamp)
	d.PhashMatrix = GreyscaleDctMatrix(greyscaleStamp)

	return nil
}
