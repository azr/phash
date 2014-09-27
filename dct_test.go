package phash

import (
	"io/ioutil"
	"testing"
	// "time"

	"fmt"
	// "github.com/azer-/phash"
	cphash "github.com/kavu/go-phash"
	"image"
	"os"
	// "code.google.com/p/biogo.matrix"
	// "github.com/hawx/img/greyscale"
	// "image/color"

	// Package image/[jpeg|fig|png] is not used explicitly in the code below,
	// but is imported for its initialization side-effect, which allows
	// image.Decode to understand [jpeg|gif|png] formatted images.
	_ "code.google.com/p/go.image/bmp"
	_ "code.google.com/p/go.image/tiff"
	_ "code.google.com/p/graphics-go/graphics"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
)

var cats_dir = "./testdata/cats/"
var lena_dir = "./testdata/lena/"

var gimages []ImageBag = nil

type Angle float64

type ImageBag struct {
	Digest
	CPhash    uint64
	Angle     Angle
	Rotations map[Angle]*ImageBag
	Dir       string
	Filename  string
	parsed    bool
}

var angles = []Angle{90, 180, 360}

var Treshold = uint64(17)

func parseDirs(ch chan<- ImageBag, dirs ...string) (images []ImageBag) {

	for _, dir := range dirs {
		done := make(chan bool)
		go func() {
			files, err := ioutil.ReadDir(dir)
			if err != nil {
				panic(err)
			}

			for _, fi := range files {
				image := getImgBag(dir, fi.Name(), 0)
				if ch == nil {
					images = append(images, *image)
				} else {
					ch <- *image
				}

				for _, angle := range angles {
					rImg := getImgBag(dir, fi.Name(), angle)
					if ch == nil {
						images = append(images, *rImg)
					} else {
						ch <- *rImg
					}
				}
			}
			done <- true
		}()
		<-done
	}
	if ch != nil {
		close(ch)
	}

	return
}

func loadImages() []ImageBag {
	if gimages == nil {
		gimages = parseDirs(nil, lena_dir, cats_dir)
	}
	return gimages
}

func loadImagesAsync() <-chan ImageBag {
	if gimages != nil {
		ch := make(chan ImageBag, len(gimages))
		for _, img := range gimages {
			ch <- img
		}
		close(ch)
		return ch
	} else {
		ch := make(chan ImageBag)
		resChan := make(chan ImageBag)
		go parseDirs(ch, lena_dir, cats_dir)
		go func() {
			var images []ImageBag
			for img := range ch {
				resChan <- img
				images = append(images, img)
			}
			close(resChan)
			gimages = images
		}()
		return resChan
	}
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
			img.ComputeGreyscaleDct(true)
		}
	}

}

func BenchmarkDctMatrix(b *testing.B) {

	images := loadImages()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		for _, img := range images {
			img.ComputeGreyscaleDctMatrix(true)
		}
	}

}

func BenchmarkDctCPhash(b *testing.B) {

	images := loadImages()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		for _, img := range images {
			img.ComputeImageHashPhash(true)
		}
	}

}

func BenchmarkRadon(b *testing.B) {

	images := loadImages()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		for _, img := range images {
			img.ComputeImageHashRadon(true)
		}
	}

}

func (img *ImageBag) ComputeGreyscaleDct(force bool) {
	if force == false && img.Digest.Phash != 0 {
		return
	}

	img.Digest.ComputeGreyscaleDct()
}

func (img *ImageBag) ComputeGreyscaleDctMatrix(force bool) {
	if force == false && img.PhashMatrix != 0 {
		return
	}

	img.Digest.ComputeGreyscaleDctMatrix()
}

func (img *ImageBag) ComputeImageHashPhash(force bool) {
	if force == false && img.CPhash != 0 {
		return
	}

	hash, err := cphash.ImageHashDCT(img.Dir + img.Filename)
	if err != nil {
		fmt.Println("Error in ComputeImageHashPhash : ", err)
	}
	img.CPhash = hash
}

func (img *ImageBag) ComputeImageHashRadon(force bool) {
	if force == false && len(img.Digest.RadonDigest.Coeffs) != 0 {
		return
	}

	// stamp := resize.Resize(32, 32, img.Image, resize.Bilinear)
	greyscaleStamp := Gscl(img.Image)
	imgMtx, err := grayImageToMatrix(greyscaleStamp)
	if err != nil {
		panic(err)
	}

	img.Digest.RadonDigest = Radon(imgMtx)
}

func (img *ImageBag) InitialiseFromFileInfo() {
	imgFile, err := os.Open(img.Dir + img.Filename) // For read access.
	defer imgFile.Close()
	if err != nil {
		panic(err)
	}
	Image, format, err := image.Decode(imgFile)
	if err != nil {
		panic(err)
	}
	img.Image = Image
	img.Format = format

	return
}

func (image *ImageBag) CompareWithImages(images []ImageBag) {
	for _, comparedImage := range images {
		if comparedImage.parsed == true {
			continue
		}

		// dist, err := cphash.HammingDistanceForHashes(image.CPhash, comparedImage.CPhash)
		// if err != nil {
		// 	dist = -1
		// }
		// dPhash1 := HammingDistance(image.PhashMatrix, comparedImage.PhashMatrix)
		// dPhash0 := HammingDistance(image.Phash, comparedImage.Phash)
		// cc := CrossCorr(image.Digest, comparedImage.Digest, 0.85)

		// fmt.Println(
		// 	"d(Phash1) : ", dPhash1,
		// 	"d(Phash0) : ", dPhash0,
		// 	" Radon : ", cc,
		// 	"c(Phash) : ", dist,
		// 	" ", image.Filename, " <> ", comparedImage.Filename)

		// if comparedImage.Filename == image.Filename {
		// 	if cc != true || dPhash0 != 0 || dPhash1 != 0 {
		// 		fmt.Println(" FAIL !")
		// 	}
		// }
	}

	return
}
