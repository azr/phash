package main

import (
"fmt"
"io/ioutil"
"os"
"image"
"image/color"
"github.com/kavu/go-phash"
    // "github.com/hawx/img/greyscale"

    // Package image/[jpeg|fig|png] is not used explicitly in the code below,
    // but is imported for its initialization side-effect, which allows
    // image.Decode to understand [jpeg|gif|png] formatted images.
_ "image/gif"
_ "image/png"
_ "image/jpeg"

"github.com/nfnt/resize"
"code.google.com/p/biogo.matrix"
"math"
"sort"
)

var images_dir = "./tests/cats/"

type ImageBag struct {
    Path string
    Format string
    CPhash uint64
    Phash0 uint64
    Phash1 uint64
}

func (image *ImageBag) CompareWithImages(images []ImageBag) {
    for _, comparedImage := range images {
        dist, err := phash.HammingDistanceForHashes(image.CPhash, comparedImage.CPhash)
        if err != nil {
            panic(err)
        }

        if (image.Path != comparedImage.Path) {
            fmt.Println( "Distance from ", image.Path, " to ", comparedImage.Path, " is : ", dist, " d(Phash0) : ", hammingDistance(image.Phash0, comparedImage.Phash0), "d(Phash1) : ", hammingDistance(image.Phash1, comparedImage.Phash1))
        }

    }

    return
}

func gscl(src image.Image) image.Gray {
    // Create a new grayscale image
    bounds := src.Bounds()
    gray := image.NewGray(bounds)
    for x := 0; x < bounds.Max.X; x++ {
        for y := 0; y < bounds.Max.Y; y++ {
            oldColor := src.At(x, y)
            gray.Set( x, y, color.GrayModel.Convert(oldColor) )
        }
    }
    return *gray
}

func coefficient(n int) (float64) {
    if (n == 0) {
        return 1.0/math.Sqrt(2)
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

            sum += math.Cos( math.Pi / ( float64(N) ) * ( float64(i) + 1/2 ) * float64(u) ) *
            math.Cos( math.Pi / ( float64(M) ) * ( float64(j) + 1/2 ) * float64(v) ) *
            float64(b)
        }
    }
    return sum * ((coefficient(u) * coefficient(v)) / 4.0)
}

func dctImageHashOne(img image.Gray) uint64 {
// func dctImageHashOne(img image.Image) ([][]float64) {
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
    avg := total / float64( ((N/2) * (M/2)) - 1);
    var hash uint64 = 0
    for u := 0; u < N/2; u++ {
        for v := 0; v < M/2; v++ {
            hash = hash * 2
            if (DCTMatrix[u][v] > avg) {
                hash++
            }
        }
    }
    return hash
}

func getImageMatrix(src image.Gray) (*matrix.Dense, error) {
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

func dctMatrix(N, M int) (*matrix.Dense, error) {
    mtx := make([][]float64, N)
    c1 := math.Sqrt(2.0/float64(N));
    c0 := 1/math.Sqrt(float64(M))
    for x := 0; x < N; x++{
        mtx[x] = make([]float64, M)
        mtx[x][0] = c0
        for y := 1; y < M; y++ {
            mtx[x][y] = c1 * math.Cos( ( math.Pi / 2.0 / float64(N) ) * float64(y) * ( 2.0 * float64(x) + 1.0 ) );
        }
    }
    return matrix.NewDense(mtx);
}

func cropMatrix(src matrix.Matrix, x0, y0, x1, y1 int) (*matrix.Dense, error) {
    mtx := make([][]float64, x1 - x0 + 1)
    for x := x0; x <= x1; x++ {
        mtx[x] = make([]float64, y1 - y0 + 1)
        for y := y0; y <= y1; y++ {
            mtx[x][y] = src.At(x, y)
        }
    }
    return matrix.NewDense(mtx)
}

func median(a []float64) float64 {
    b := make([]float64, len(a))
    copy(b, a)
    sort.Float64s(b)
    return b[int(len(b) / 2)]
}

func dctImageHashTwo(img image.Gray) uint64 {
    imgMtx, err := getImageMatrix(img)
    if err != nil {
        panic(err)
    }
    dctMtx, err  := dctMatrix(img.Bounds().Max.X, img.Bounds().Max.Y)
    if err != nil {
        panic(err)
    }
    dctMtxTransp := dctMtx.T(nil); // Transpose

    // CImg<float> dctImage = (*C)*img*Ctransp;
    dctImage := dctMtx.Dot( imgMtx, nil ).Dot(dctMtxTransp, nil)

    dctImage, err = cropMatrix(dctImage, 0, 0, 7, 7)
    if err != nil {
        panic(err)
    }
    subsec := matrix.ElementsVector(dctImage)
    median := median(subsec)
    var one , hash uint64 = 1, 0
    for  i := 0; i < len(subsec); i++ {
        current := subsec[i]
        if (current > median) {
            hash |= one
        }
        one = one << 1
    }
    return hash
}

func hammingDistance( hash1, hash2 uint64) uint64 {
    x := hash1^hash2;
    var m1, m2, h01, m4 uint64 = 0x5555555555555555, 0x3333333333333333, 0x0101010101010101, 0x0f0f0f0f0f0f0f0f
    x -= (x >> 1) & m1;
    x = (x & m2) + ((x >> 2) & m2);
    x = (x + (x >> 4)) & m4;
    return (x * h01)>>56;
}

// 32x32 gray image
func (img *ImageBag) ComputeImageHash() {
    imgFile, err := os.Open(img.Path) // For read access.
    if err != nil {
        panic(err)
    }
    decodedImage, format, err := image.Decode(imgFile)
    if err != nil {
        panic(err)
    }
    imgFile.Close()
    img.Format = format

    stamp := resize.Resize(32, 32, decodedImage, resize.Bilinear)
    greyscaleStamp := gscl(stamp)
    // greyscaleStamp := greyscale.Greyscale(stamp)
    img.Phash0 = dctImageHashOne(greyscaleStamp)
    img.Phash1 = dctImageHashTwo(greyscaleStamp)

    fmt.Println( "CPhash : ", img.CPhash, "Phash0 : ", img.Phash0, " Phash1 : ", img.Phash1, " for : ", img.Path)
}

func (img *ImageBag) InitialiseFromFileInfo(fi os.FileInfo) {
    img.Path = images_dir + fi.Name()
    hash, err := phash.ImageHashDCT(img.Path)
    if err != nil {
        panic(err)
    }
    img.CPhash = hash
    img.ComputeImageHash()

    return
}

func main() {

    images_in_dir, err := ioutil.ReadDir(images_dir)
    if err != nil {
        panic(err)
    }

    images := make([]ImageBag, len(images_in_dir) )
    for i, fi := range images_in_dir {
        images[i].InitialiseFromFileInfo(fi)
    }

    fmt.Println( "Loaded images" )
    for _, img := range images {
        img.CompareWithImages(images)
    }

}

