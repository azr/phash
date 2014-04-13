package phash

import (
	"io/ioutil"
	"testing"
	// "time"

	"fmt"
	"image"
	"os"
	cphash "github.com/kavu/go-phash"
	// "code.google.com/p/biogo.matrix"
	// "github.com/hawx/img/greyscale"
	// "image/color"

	// Package image/[jpeg|fig|png] is not used explicitly in the code below,
	// but is imported for its initialization side-effect, which allows
	// image.Decode to understand [jpeg|gif|png] formatted images.
	_ "code.google.com/p/go.image/bmp"
	_ "code.google.com/p/go.image/tiff"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"image/draw"

	"github.com/nfnt/resize"
	"code.google.com/p/graphics-go/graphics"
)

var cats_dir = "./testdata/cats/"
var lena_dir = "./testdata/lena/"

var gimages []ImageBag = nil

type ImageBag struct {
	Dir          string
	Path         string
	decodedImage image.Image
	Format       string
	Angle        float64
	CPhash       uint64
	Phash0       uint64
	Phash1       uint64
	digest       Digest
	parsed       bool
}

var angles = []float64 { 90, 180, 360 }

func copyImage(src image.Image) (draw.Image) {
	b := src.Bounds()
	copy := image.NewRGBA(b)

	draw.Draw(copy, copy.Bounds(), src, b.Min, draw.Src)

	return copy
}

func parseDirs(dirs ...string) (images []ImageBag) {
	for _, dir := range dirs {
		files_in_dir, err := ioutil.ReadDir(dir)
		if err != nil {
			panic(err)
		}
		for _, fi := range files_in_dir {
			image := ImageBag{dir, dir + fi.Name(), nil, "", 0, 0, 0, 0, Digest{}, false}
			image.InitialiseFromFileInfo()
			images = append(images, image)
			for angle := range angles {

				newImg := copyImage(image.decodedImage)

				if err := graphics.Rotate(newImg, image.decodedImage, &graphics.RotateOptions{float64(angle)}) ; err != nil {
					panic( err )
				}
				rImg := ImageBag{dir, dir + fi.Name(), newImg, image.Format, float64(angle), 0, 0, 0, Digest{}, false}
				images = append(images, rImg)
			}
		}
	}

	return
}

func loadImages() []ImageBag {
	if gimages == nil {
		gimages = parseDirs(lena_dir, cats_dir)
	}
	return gimages
}

// func TestTimeConsuming(t *testing.T) {
//     if testing.Short() {
//         t.Skip("skipping test in short mode.")
//     }

// }

func BenchmarkDct(b *testing.B) {

	images := loadImages()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		for _, img := range images {
			img.ComputeDct()
		}
	}

}

func BenchmarkDctMatrix(b *testing.B) {

	images := loadImages()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		for _, img := range images {
			img.ComputeDctMatrix()
		}
	}

}

func BenchmarkDctCPhash(b *testing.B) {

	images := loadImages()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		for _, img := range images {
			img.ComputeImageHashPhash()
		}
	}

}

func BenchmarkRadon(b *testing.B) {

	images := loadImages()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		for _, img := range images {
			img.ComputeImageHashRadon()
		}
	}

}

func (img *ImageBag) ComputeDct() {

	stamp := resize.Resize(32, 32, img.decodedImage, resize.Bilinear)
	greyscaleStamp := gscl(stamp)

	// greyscaleStamp := greyscale.Greyscale(stamp)
	img.Phash0 = Dct(greyscaleStamp)
}

func (img *ImageBag) ComputeDctMatrix() {

	stamp := resize.Resize(32, 32, img.decodedImage, resize.Bilinear)
	greyscaleStamp := gscl(stamp)
	// greyscaleStamp := greyscale.Greyscale(stamp)
	img.Phash1 = DctMatrix(greyscaleStamp)
}

func (img *ImageBag) ComputeImageHashPhash() {
	hash, err := cphash.ImageHashDCT(img.Path)
	if err != nil {
		fmt.Println("Error in ComputeImageHashPhash : ", err)
	}
	img.CPhash = hash
}

func (img *ImageBag) ComputeImageHashRadon() {
	// stamp := resize.Resize(32, 32, img.decodedImage, resize.Bilinear)
	greyscaleStamp := gscl(img.decodedImage)
	img.digest = Radon(greyscaleStamp)
}

func (img *ImageBag) InitialiseFromFileInfo() {
	imgFile, err := os.Open(img.Path) // For read access.
	defer imgFile.Close()
	if err != nil {
		panic(err)
	}
	decodedImage, format, err := image.Decode(imgFile)
	if err != nil {
		panic(err)
	}
	img.decodedImage = decodedImage
	img.Format = format

	return
}

func (image *ImageBag) CompareWithImages(images []ImageBag) {
	for _, comparedImage := range images {
		if comparedImage.parsed == true {
			continue
		}

		dist, err := cphash.HammingDistanceForHashes(image.CPhash, comparedImage.CPhash)
		if err != nil {
			dist = -1
		}
		dPhash1 := hammingDistance(image.Phash1, comparedImage.Phash1)
		dPhash0 := hammingDistance(image.Phash0, comparedImage.Phash0)
		cc := crosscorr(image.digest, comparedImage.digest, 0.85)

		fmt.Println(
			"d(Phash1) : ", dPhash1,
			"d(Phash0) : ", dPhash0,
			" Radon : ", cc,
			"c(Phash) : ", dist,
			" ", image.Path, " <> ", comparedImage.Path)

		if comparedImage.Path == image.Path {
			if cc != true || dPhash0 != 0 || dPhash1 != 0 {
				fmt.Println(" LOL !")
				//panic(" I'm an asshole with " + image.Path)
			}
		}

	}

	return
}
