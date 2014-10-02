package phash

import (
	// "github.com/azer-/phash"
	"code.google.com/p/graphics-go/graphics"
	"github.com/azr/phash/radon"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

var testImages = map[string]*ImageBag{}

var Treshold = uint64(17)

func getImgBag(dir, filename string, angle Angle) *ImageBag {
	img, found := testImages[dir+filename]
	if !found {
		testImages[dir+filename] = new(ImageBag)
		testImages[dir+filename].Dir = dir
		testImages[dir+filename].Filename = filename
		testImages[dir+filename].Rotations = map[Angle]*ImageBag{}
		testImages[dir+filename].InitialiseFromFileInfo()
		img = testImages[dir+filename]
	}

	if angle != 0 {
		rotatedImage, found := img.Rotations[angle]
		if !found {
			draw := CopyImage(img.Image)
			err := graphics.Rotate(draw, img.Image, &graphics.RotateOptions{Angle: float64(angle)})
			if err != nil {
				panic(err)
			}
			rImg := ImageBag{Dir: dir, Filename: filename, Angle: angle, ImageDigest: radon.ImageDigest{Image: draw, Format: img.Format}}
			img.Rotations[angle] = &rImg

			return img.Rotations[angle]
		}
		return rotatedImage
	}

	return img
}

func PairExecuteFor2ImagesList(imagesA []*ImageBag, imagesB []*ImageBag, f func(*ImageBag, *ImageBag)) {
	for _, imageA := range imagesA {
		for _, imageB := range imagesB {
			f(imageA, imageB)
		}
	}
}

func TestAffineTransformedImagesMatch(t *testing.T) {

	Convey("Given some cats are loaded", t, func() {
		catBig := getImgBag(catsDir, "cat_big.jpg", 0)
		catMedium := getImgBag(catsDir, "cat_medium.jpg", 0)
		catSmall := getImgBag(catsDir, "cat_small.jpg", 0)

		Convey("When CDCT Hashes are computed", func() {
			catBig.ComputeImageHashPhash(false)
			catMedium.ComputeImageHashPhash(false)
			catSmall.ComputeImageHashPhash(false)

			Convey("Then the hashes should not be zero", func() {
				So(catBig.CPhash, ShouldNotEqual, 0)
				So(catMedium.CPhash, ShouldNotEqual, 0)
				So(catSmall.CPhash, ShouldNotEqual, 0)
			})

			Convey("Then Hamming Distance should be under threshold ", func() {
				distance := HammingDistance(catBig.CPhash, catMedium.CPhash)
				So(distance, ShouldBeLessThanOrEqualTo, Treshold)
				distance = HammingDistance(catBig.CPhash, catSmall.CPhash)
				So(distance, ShouldBeLessThanOrEqualTo, Treshold)
				distance = HammingDistance(catMedium.CPhash, catSmall.CPhash)
				So(distance, ShouldBeLessThanOrEqualTo, Treshold)
			})

		})

		SkipConvey("When Radon Hashes are computed", func() {
			catBig.ComputeImageHashRadon(false)
			catMedium.ComputeImageHashRadon(false)

			Convey("Then the hashes should not be zero", func() {
				So(len(catBig.ImageDigest.ImageDigest.Coeffs), ShouldNotEqual, 0)
				So(catBig.ImageDigest.ImageDigest.Coeffs, ShouldNotBeNil)
				So(len(catMedium.ImageDigest.ImageDigest.Coeffs), ShouldNotEqual, 0)
				So(catMedium.ImageDigest.ImageDigest.Coeffs, ShouldNotBeNil)
			})

			Convey("Then the Cross Correlation for threshold should be true", func() {
				crossCoorrelation, _ := radon.CrossCorr(catBig.ImageDigest.ImageDigest, catMedium.ImageDigest.ImageDigest, 0.01)
				So(crossCoorrelation, ShouldBeTrue)
			})

		})

	})

	Convey("Given some suns are loaded", t, func() {
		sunBig := getImgBag(catsDir, "sun_big.jpg", 0)
		sunSmall := getImgBag(catsDir, "sun_small.jpg", 0)

		Convey("When CDCT Hashes are computed", func() {
			sunBig.ComputeImageHashPhash(false)
			sunSmall.ComputeImageHashPhash(false)

			Convey("Then the hashes should not be zero", func() {
				So(sunBig.CPhash, ShouldNotEqual, 0)
				So(sunSmall.CPhash, ShouldNotEqual, 0)
			})

			Convey("Then Hamming Distance should be under threshold ", func() {
				distance := HammingDistance(sunBig.CPhash, sunSmall.CPhash)
				So(distance, ShouldBeLessThanOrEqualTo, Treshold)
				distance = HammingDistance(sunBig.CPhash, sunSmall.CPhash)
				So(distance, ShouldBeLessThanOrEqualTo, Treshold)
			})

		})

		Convey("When Radon Hashes are computed", func() {
			sunBig.ComputeImageHashRadon(false)
			sunSmall.ComputeImageHashRadon(false)

			Convey("Then the hashes should not be zero", func() {
				So(len(sunBig.ImageDigest.ImageDigest.Coeffs), ShouldNotEqual, 0)
				So(sunBig.ImageDigest.ImageDigest.Coeffs, ShouldNotBeNil)
				So(len(sunSmall.ImageDigest.ImageDigest.Coeffs), ShouldNotEqual, 0)
				So(sunSmall.ImageDigest.ImageDigest.Coeffs, ShouldNotBeNil)
			})

			SkipConvey("Then the Cross Correlation for threshold should be true", func() {

				crossCoorrelation, _ := radon.CrossCorr(sunBig.ImageDigest.ImageDigest, sunSmall.ImageDigest.ImageDigest, -0.1)
				So(crossCoorrelation, ShouldBeTrue)
				crossCoorrelation, _ = radon.CrossCorr(sunBig.ImageDigest.ImageDigest, sunSmall.ImageDigest.ImageDigest, 0.0)
				So(crossCoorrelation, ShouldBeTrue)
			})

		})

	})

}

func TestDifferentImagesDoNotMatch(t *testing.T) {

	Convey("Given some cats, suns and girls are loaded", t, func() {

		catBig := getImgBag(catsDir, "cat_big.jpg", 0)
		catMedium := getImgBag(catsDir, "cat_medium.jpg", 0)
		catSmall := getImgBag(catsDir, "cat_small.jpg", 0)
		catSky := getImgBag(catsDir, "cat_sky.jpg", 0)
		catSmile := getImgBag(catsDir, "smiling.jpg", 0)
		cats := []*ImageBag{catBig, catMedium, catSmall, catSky, catSmile}

		lenaToshop := getImgBag(lenaDir, "lena_std.tiff", 0)
		lenaNkd := getImgBag(lenaDir, "l_hires.jpg", 0)
		lenaVintage := getImgBag(lenaDir, "lena.bmp", 0)
		girls := []*ImageBag{lenaToshop, lenaNkd, lenaVintage}

		sunBig := getImgBag(catsDir, "sun_big.jpg", 0)
		sunSmall := getImgBag(catsDir, "sun_small.jpg", 0)
		suns := []*ImageBag{sunBig, sunSmall}

		Convey("When CDCT Hashes are computed", func() {

			catBig.ComputeImageHashPhash(false)
			catMedium.ComputeImageHashPhash(false)
			catSmall.ComputeImageHashPhash(false)
			catSky.ComputeImageHashPhash(false)
			catSmile.ComputeImageHashPhash(false)

			// lenaToshop.ComputeImageHashPhash(false)
			lenaVintage.ComputeImageHashPhash(false)
			lenaNkd.ComputeImageHashPhash(false)
			girls := []*ImageBag{lenaVintage, lenaNkd}

			sunBig.ComputeImageHashPhash(false)
			sunSmall.ComputeImageHashPhash(false)

			Convey("Then cats do not look like girls", func() {
				PairExecuteFor2ImagesList(cats, girls, func(cat, girl *ImageBag) {
					distance := HammingDistance(cat.CPhash, girl.CPhash)
					So(distance, ShouldBeGreaterThan, Treshold)
				})
			})

			Convey("And suns do not look like cats", func() {
				PairExecuteFor2ImagesList(suns, cats, func(sun, cat *ImageBag) {
					distance := HammingDistance(sun.CPhash, cat.CPhash)
					So(distance, ShouldBeGreaterThan, Treshold)
				})
			})

			Convey("And girls do not look like suns", func() {
				PairExecuteFor2ImagesList(girls, suns, func(girl, sun *ImageBag) {
					distance := HammingDistance(sun.CPhash, girl.CPhash)
					So(distance, ShouldBeGreaterThan, Treshold)
				})
			})

		})

		SkipConvey("When Radon Hashes are computed", func() {

			catBig.ComputeImageHashRadon(false)
			catMedium.ComputeImageHashRadon(false)
			catSmall.ComputeImageHashRadon(false)
			catSky.ComputeImageHashRadon(false)
			catSmile.ComputeImageHashRadon(false)

			lenaToshop.ComputeImageHashRadon(false)
			lenaNkd.ComputeImageHashRadon(false)
			lenaVintage.ComputeImageHashRadon(false)

			sunBig.ComputeImageHashRadon(false)
			sunSmall.ComputeImageHashRadon(false)

			Convey("Then cats do not look like girls", func() {
				PairExecuteFor2ImagesList(cats, girls, func(cat, girl *ImageBag) {
					crossCoorrelation, _ := radon.CrossCorr(cat.ImageDigest.ImageDigest, girl.ImageDigest.ImageDigest, 0.8)
					So(crossCoorrelation, ShouldBeFalse)
				})
			})

			Convey("And suns do not look like cats", func() {
				PairExecuteFor2ImagesList(suns, cats, func(sun, cat *ImageBag) {
					crossCoorrelation, _ := radon.CrossCorr(sun.ImageDigest.ImageDigest, cat.ImageDigest.ImageDigest, 0.8)
					So(crossCoorrelation, ShouldBeFalse)
				})
			})

			Convey("And girls do not look like suns", func() {
				PairExecuteFor2ImagesList(girls, suns, func(girl, sun *ImageBag) {
					crossCoorrelation, _ := radon.CrossCorr(girl.ImageDigest.ImageDigest, sun.ImageDigest.ImageDigest, 0.8)
					So(crossCoorrelation, ShouldBeFalse)
				})
			})

		})

	})

}
