package phash

import (
	// "github.com/azer-/phash"
	"code.google.com/p/graphics-go/graphics"
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
			rImg := ImageBag{Dir: dir, Filename: filename, Angle: angle, Digest: Digest{Image: draw, Format: img.Format}}
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

		Convey("When DCT Hashes are computed", func() {
			catBig.ComputeGreyscaleDct(false)
			catMedium.ComputeGreyscaleDct(false)
			catSmall.ComputeGreyscaleDct(false)

			Convey("Then the hashes should not be zero", func() {
				So(catBig.Phash, ShouldNotEqual, 0)
				So(catMedium.Phash, ShouldNotEqual, 0)
				So(catSmall.Phash, ShouldNotEqual, 0)
			})

			Convey("And Hamming Distance should be under threshold ", func() {
				distance := HammingDistance(catBig.Phash, catMedium.Phash)
				So(distance, ShouldBeLessThanOrEqualTo, Treshold)
				distance = HammingDistance(catBig.Phash, catSmall.Phash)
				So(distance, ShouldBeLessThanOrEqualTo, Treshold)
				distance = HammingDistance(catMedium.Phash, catSmall.Phash)
				So(distance, ShouldBeLessThanOrEqualTo, Treshold)
			})

		})

		Convey("When DCTMatrix Hashes are computed", func() {
			catBig.ComputeGreyscaleDctMatrix(false)
			catMedium.ComputeGreyscaleDctMatrix(false)
			catSmall.ComputeGreyscaleDctMatrix(false)

			Convey("Then the hashes should not be zero", func() {
				So(catBig.PhashMatrix, ShouldNotEqual, 0)
				So(catMedium.PhashMatrix, ShouldNotEqual, 0)
				So(catSmall.PhashMatrix, ShouldNotEqual, 0)
			})

			Convey("Then Hamming Distance should be under threshold ", func() {
				distance := HammingDistance(catBig.PhashMatrix, catMedium.PhashMatrix)
				So(distance, ShouldBeLessThanOrEqualTo, Treshold)
				distance = HammingDistance(catBig.PhashMatrix, catSmall.PhashMatrix)
				So(distance, ShouldBeLessThanOrEqualTo, Treshold)
				distance = HammingDistance(catMedium.PhashMatrix, catSmall.PhashMatrix)
				So(distance, ShouldBeLessThanOrEqualTo, Treshold)
			})

		})

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
				So(len(catBig.Digest.RadonDigest.Coeffs), ShouldNotEqual, 0)
				So(catBig.Digest.RadonDigest.Coeffs, ShouldNotBeNil)
				So(len(catMedium.Digest.RadonDigest.Coeffs), ShouldNotEqual, 0)
				So(catMedium.Digest.RadonDigest.Coeffs, ShouldNotBeNil)
			})

			Convey("Then the Cross Correlation for threshold should be true", func() {
				crossCoorrelation, _ := CrossCorr(catBig.Digest.RadonDigest, catMedium.Digest.RadonDigest, 0.01)
				So(crossCoorrelation, ShouldBeTrue)
			})

		})

	})

	Convey("Given some suns are loaded", t, func() {
		sunBig := getImgBag(catsDir, "sun_big.jpg", 0)
		sunSmall := getImgBag(catsDir, "sun_small.jpg", 0)

		Convey("When DCT Hashes are computed", func() {
			sunBig.ComputeGreyscaleDct(false)
			sunSmall.ComputeGreyscaleDct(false)

			Convey("Then the hashes should not be zero", func() {
				So(sunBig.Phash, ShouldNotEqual, 0)
				So(sunSmall.Phash, ShouldNotEqual, 0)
			})

			Convey("And Hamming Distance should be under threshold ", func() {
				distance := HammingDistance(sunBig.Phash, sunSmall.Phash)
				So(distance, ShouldBeLessThanOrEqualTo, Treshold)
			})

		})

		Convey("When DCTMatrix Hashes are computed", func() {
			sunBig.ComputeGreyscaleDctMatrix(false)
			sunSmall.ComputeGreyscaleDctMatrix(false)

			Convey("Then the hashes should not be zero", func() {
				So(sunBig.PhashMatrix, ShouldNotEqual, 0)
				So(sunSmall.PhashMatrix, ShouldNotEqual, 0)
			})

			Convey("Then Hamming Distance should be under threshold ", func() {
				distance := HammingDistance(sunBig.PhashMatrix, sunSmall.PhashMatrix)
				So(distance, ShouldBeLessThanOrEqualTo, Treshold)
				distance = HammingDistance(sunBig.PhashMatrix, sunSmall.PhashMatrix)
				So(distance, ShouldBeLessThanOrEqualTo, Treshold)
			})

		})

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
				So(len(sunBig.Digest.RadonDigest.Coeffs), ShouldNotEqual, 0)
				So(sunBig.Digest.RadonDigest.Coeffs, ShouldNotBeNil)
				So(len(sunSmall.Digest.RadonDigest.Coeffs), ShouldNotEqual, 0)
				So(sunSmall.Digest.RadonDigest.Coeffs, ShouldNotBeNil)
			})

			SkipConvey("Then the Cross Correlation for threshold should be true", func() {

				crossCoorrelation, _ := CrossCorr(sunBig.Digest.RadonDigest, sunSmall.Digest.RadonDigest, -0.1)
				So(crossCoorrelation, ShouldBeTrue)
				crossCoorrelation, _ = CrossCorr(sunBig.Digest.RadonDigest, sunSmall.Digest.RadonDigest, 0.0)
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

		Convey("When DCT Hashes are computed", func() {

			catBig.ComputeGreyscaleDct(false)
			catMedium.ComputeGreyscaleDct(false)
			catSmall.ComputeGreyscaleDct(false)
			catSky.ComputeGreyscaleDct(false)
			catSmile.ComputeGreyscaleDct(false)

			lenaToshop.ComputeGreyscaleDct(false)
			lenaNkd.ComputeGreyscaleDct(false)
			lenaVintage.ComputeGreyscaleDct(false)

			sunBig.ComputeGreyscaleDct(false)
			sunSmall.ComputeGreyscaleDct(false)

			Convey("Then cats do not look like girls", func() {
				PairExecuteFor2ImagesList(cats, girls, func(cat, girl *ImageBag) {
					distance := HammingDistance(cat.Phash, girl.Phash)
					So(distance, ShouldBeGreaterThan, Treshold)
				})
			})

			Convey("And suns do not look like cats", func() {
				PairExecuteFor2ImagesList(suns, cats, func(sun, cat *ImageBag) {
					distance := HammingDistance(sun.Phash, cat.Phash)
					So(distance, ShouldBeGreaterThan, Treshold)
				})
			})

			Convey("And girls do not look like suns", func() {
				PairExecuteFor2ImagesList(girls, suns, func(girl, sun *ImageBag) {
					distance := HammingDistance(sun.Phash, girl.Phash)
					So(distance, ShouldBeGreaterThan, Treshold)
				})
			})

		})

		Convey("When DCTMatrix Hashes are computed", func() {

			catBig.ComputeGreyscaleDctMatrix(false)
			catMedium.ComputeGreyscaleDctMatrix(false)
			catSmall.ComputeGreyscaleDctMatrix(false)
			catSky.ComputeGreyscaleDctMatrix(false)
			catSmile.ComputeGreyscaleDctMatrix(false)

			lenaToshop.ComputeGreyscaleDctMatrix(false)
			lenaNkd.ComputeGreyscaleDctMatrix(false)
			lenaVintage.ComputeGreyscaleDctMatrix(false)

			sunBig.ComputeGreyscaleDctMatrix(false)
			sunSmall.ComputeGreyscaleDctMatrix(false)

			Convey("Then cats do not look like girls", func() {
				PairExecuteFor2ImagesList(cats, girls, func(cat, girl *ImageBag) {
					distance := HammingDistance(cat.PhashMatrix, girl.PhashMatrix)
					So(distance, ShouldBeGreaterThan, Treshold)
				})
			})

			Convey("And suns do not look like cats", func() {
				PairExecuteFor2ImagesList(suns, cats, func(sun, cat *ImageBag) {
					distance := HammingDistance(sun.PhashMatrix, cat.PhashMatrix)
					So(distance, ShouldBeGreaterThan, Treshold)
				})
			})

			Convey("And girls do not look like suns", func() {
				PairExecuteFor2ImagesList(girls, suns, func(girl, sun *ImageBag) {
					distance := HammingDistance(sun.PhashMatrix, girl.PhashMatrix)
					So(distance, ShouldBeGreaterThan, Treshold)
				})
			})

		})

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
					crossCoorrelation, _ := CrossCorr(cat.Digest.RadonDigest, girl.Digest.RadonDigest, 0.8)
					So(crossCoorrelation, ShouldBeFalse)
				})
			})

			Convey("And suns do not look like cats", func() {
				PairExecuteFor2ImagesList(suns, cats, func(sun, cat *ImageBag) {
					crossCoorrelation, _ := CrossCorr(sun.Digest.RadonDigest, cat.Digest.RadonDigest, 0.8)
					So(crossCoorrelation, ShouldBeFalse)
				})
			})

			Convey("And girls do not look like suns", func() {
				PairExecuteFor2ImagesList(girls, suns, func(girl, sun *ImageBag) {
					crossCoorrelation, _ := CrossCorr(girl.Digest.RadonDigest, sun.Digest.RadonDigest, 0.8)
					So(crossCoorrelation, ShouldBeFalse)
				})
			})

		})

	})

}
