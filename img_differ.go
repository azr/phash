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
_ "code.google.com/p/go.image/tiff"
_ "code.google.com/p/go.image/bmp"

"github.com/nfnt/resize"
)

var cats_dir = "./tests/cats/"
var lena_dir = "./tests/lena/"

func (image *ImageBag) CompareWithImages(images []ImageBag) {
    for _, comparedImage := range images {
        if (comparedImage.parsed == true) {
            continue
        }

        dist, err := phash.HammingDistanceForHashes(image.CPhash, comparedImage.CPhash)
        if err != nil {
            dist = -1
        }
        dPhash1 := hammingDistance(image.Phash1, comparedImage.Phash1)
        dPhash0 := hammingDistance(image.Phash0, comparedImage.Phash0)
        cc := crosscorr(image.digest, comparedImage.digest, 0.85)

        fmt.Println(
            "d(Phash1) : ", dPhash1 ,
            "d(Phash0) : ", dPhash0,
            " Radon : ", cc,
            "c(Phash) : ", dist,
            " ", image.Path, " <> " , comparedImage.Path )

        if (comparedImage.Path == image.Path) {
            if (cc != true || dPhash0 != 0 || dPhash1 != 0) {
                fmt.Println(" LOL !")
                //panic(" I'm an asshole with " + image.Path)
            }
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
        fmt.Println("Error in ComputeImageHashPhash : ", err)
    }
    img.CPhash = hash
}

func (img *ImageBag) ComputeImageHashRadon() {
    // stamp := resize.Resize(32, 32, img.decodedImage, resize.Bilinear)
    greyscaleStamp := gscl(img.decodedImage)
    img.digest = DctImageHashRadon(greyscaleStamp)
}

func (img *ImageBag) InitialiseFromFileInfo() {
    imgFile, err := os.Open(img.Path) // For read access.
    if err != nil {
        panic(err)
    }
    fmt.Println( "Open ", img.Path )
    decodedImage, format, err := image.Decode(imgFile)
    if err != nil {
        panic(err)
    }
    imgFile.Close()
    img.decodedImage = decodedImage
    img.Format = format

    return
}

func parseDirs(dirs ...string) []ImageBag {
    var images []ImageBag
    for _, dir := range dirs {
        files_in_dir, err := ioutil.ReadDir(dir)
        if err != nil {
            panic(err)
        }
        for _, fi := range files_in_dir {
            image := ImageBag{dir, dir + fi.Name(), nil, "", 0, 0, 0, Digest{}, false}
            image.InitialiseFromFileInfo()
            image.ComputeImageHashOne()
            image.ComputeImageHashTwo()
            image.ComputeImageHashPhash()
            image.ComputeImageHashRadon()
            images = append(images, image)
        }
    }

    return images
}

func main() {
    images := parseDirs(lena_dir, cats_dir)

    fmt.Println( "Loaded and populated images" )
    for i, img := range images {
        img.CompareWithImages(images)
        images[i].parsed = true
    }
}

