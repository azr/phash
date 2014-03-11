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
)

var images_dir = "./tests/cats/"

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

func (img *ImageBag) ComputeImageHashOne() {

    stamp := resize.Resize(32, 32, img.decodedImage, resize.Bilinear)
    greyscaleStamp := gscl(stamp)
    // greyscaleStamp := greyscale.Greyscale(stamp)
    img.Phash0 = DctImageHashOne(greyscaleStamp)
}

func (img *ImageBag) ComputeImageHashTwo() {

    stamp := resize.Resize(32, 32, img.decodedImage, resize.Bilinear)
    greyscaleStamp := gscl(stamp)
    // greyscaleStamp := greyscale.Greyscale(stamp)
    img.Phash1 = DctImageHashTwo(greyscaleStamp)
}

func (img *ImageBag) ComputeImageHashPhash() {
    hash, err := phash.ImageHashDCT(img.Path)
    if err != nil {
        panic(err)
    }
    img.CPhash = hash
}

func (img *ImageBag) InitialiseFromFileInfo(fi os.FileInfo) {
    img.Path = images_dir + fi.Name()
    imgFile, err := os.Open(img.Path) // For read access.
    if err != nil {
        panic(err)
    }
    decodedImage, format, err := image.Decode(imgFile)
    if err != nil {
        panic(err)
    }
    imgFile.Close()
    img.decodedImage = decodedImage
    img.Format = format

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
        images[i].ComputeImageHashOne()
        images[i].ComputeImageHashTwo()
        images[i].ComputeImageHashPhash()
    }

    fmt.Println( "Loaded images" )
    for _, img := range images {
        img.CompareWithImages(images)
    }

}

