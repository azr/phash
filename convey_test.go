package phash

import (
	// "github.com/azer-/phash"
	"code.google.com/p/graphics-go/graphics"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

var testImages map[string]*ImageBag = map[string]*ImageBag{}

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
		if rotatedImage, found := img.Rotations[angle]; !found {
			draw := CopyImage(img.Image)
			if err := graphics.Rotate(draw, img.Image, &graphics.RotateOptions{float64(angle)}); err != nil {
				panic(err)
			}
			rImg := ImageBag{Dir: dir, Filename: filename, Angle: angle, Digest: Digest{Image: draw, Format: img.Format}}
			img.Rotations[angle] = &rImg

			return img.Rotations[angle]
		} else {
			return rotatedImage
		}
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
		cat_big := getImgBag(cats_dir, "cat_big.jpg", 0)
		cat_medium := getImgBag(cats_dir, "cat_medium.jpg", 0)
		cat_small := getImgBag(cats_dir, "cat_small.jpg", 0)

		Convey("When DCT Hashes are computed", func() {
			cat_big.ComputeGreyscaleDct(false)
			cat_medium.ComputeGreyscaleDct(false)
			cat_small.ComputeGreyscaleDct(false)

			Convey("Then the hashes should not be zero", func() {
				So(cat_big.Phash, ShouldNotEqual, 0)
				So(cat_medium.Phash, ShouldNotEqual, 0)
				So(cat_small.Phash, ShouldNotEqual, 0)
			})

			Convey("And Hamming Distance should be under threshold ", func() {
				distance := HammingDistance(cat_big.Phash, cat_medium.Phash)
				So(distance, ShouldBeLessThanOrEqualTo, Treshold)
				distance = HammingDistance(cat_big.Phash, cat_small.Phash)
				So(distance, ShouldBeLessThanOrEqualTo, Treshold)
				distance = HammingDistance(cat_medium.Phash, cat_small.Phash)
				So(distance, ShouldBeLessThanOrEqualTo, Treshold)
			})

		})

		Convey("When DCTMatrix Hashes are computed", func() {
			cat_big.ComputeGreyscaleDctMatrix(false)
			cat_medium.ComputeGreyscaleDctMatrix(false)
			cat_small.ComputeGreyscaleDctMatrix(false)

			Convey("Then the hashes should not be zero", func() {
				So(cat_big.PhashMatrix, ShouldNotEqual, 0)
				So(cat_medium.PhashMatrix, ShouldNotEqual, 0)
				So(cat_small.PhashMatrix, ShouldNotEqual, 0)
			})

			Convey("Then Hamming Distance should be under threshold ", func() {
				distance := HammingDistance(cat_big.PhashMatrix, cat_medium.PhashMatrix)
				So(distance, ShouldBeLessThanOrEqualTo, Treshold)
				distance = HammingDistance(cat_big.PhashMatrix, cat_small.PhashMatrix)
				So(distance, ShouldBeLessThanOrEqualTo, Treshold)
				distance = HammingDistance(cat_medium.PhashMatrix, cat_small.PhashMatrix)
				So(distance, ShouldBeLessThanOrEqualTo, Treshold)
			})

		})

		Convey("When CDCT Hashes are computed", func() {
			cat_big.ComputeImageHashPhash(false)
			cat_medium.ComputeImageHashPhash(false)
			cat_small.ComputeImageHashPhash(false)

			Convey("Then the hashes should not be zero", func() {
				So(cat_big.CPhash, ShouldNotEqual, 0)
				So(cat_medium.CPhash, ShouldNotEqual, 0)
				So(cat_small.CPhash, ShouldNotEqual, 0)
			})

			Convey("Then Hamming Distance should be under threshold ", func() {
				distance := HammingDistance(cat_big.CPhash, cat_medium.CPhash)
				So(distance, ShouldBeLessThanOrEqualTo, Treshold)
				distance = HammingDistance(cat_big.CPhash, cat_small.CPhash)
				So(distance, ShouldBeLessThanOrEqualTo, Treshold)
				distance = HammingDistance(cat_medium.CPhash, cat_small.CPhash)
				So(distance, ShouldBeLessThanOrEqualTo, Treshold)
			})

		})

		SkipConvey("When Radon Hashes are computed", func() {
			cat_big.ComputeImageHashRadon(false)
			cat_medium.ComputeImageHashRadon(false)

			Convey("Then the hashes should not be zero", func() {
				So(len(cat_big.Digest.RadonDigest.Coeffs), ShouldNotEqual, 0)
				So(cat_big.Digest.RadonDigest.Coeffs, ShouldNotBeNil)
				So(len(cat_medium.Digest.RadonDigest.Coeffs), ShouldNotEqual, 0)
				So(cat_medium.Digest.RadonDigest.Coeffs, ShouldNotBeNil)
			})

			Convey("Then the Cross Correlation for threshold should be true", func() {
				So(CrossCorr(cat_big.Digest.RadonDigest, cat_medium.Digest.RadonDigest, 0.01), ShouldBeTrue)
			})

		})

	})

	Convey("Given some suns are loaded", t, func() {
		sun_big := getImgBag(cats_dir, "sun_big.jpg", 0)
		sun_small := getImgBag(cats_dir, "sun_small.jpg", 0)

		Convey("When DCT Hashes are computed", func() {
			sun_big.ComputeGreyscaleDct(false)
			sun_small.ComputeGreyscaleDct(false)

			Convey("Then the hashes should not be zero", func() {
				So(sun_big.Phash, ShouldNotEqual, 0)
				So(sun_small.Phash, ShouldNotEqual, 0)
			})

			Convey("And Hamming Distance should be under threshold ", func() {
				distance := HammingDistance(sun_big.Phash, sun_small.Phash)
				So(distance, ShouldBeLessThanOrEqualTo, Treshold)
			})

		})

		Convey("When DCTMatrix Hashes are computed", func() {
			sun_big.ComputeGreyscaleDctMatrix(false)
			sun_small.ComputeGreyscaleDctMatrix(false)

			Convey("Then the hashes should not be zero", func() {
				So(sun_big.PhashMatrix, ShouldNotEqual, 0)
				So(sun_small.PhashMatrix, ShouldNotEqual, 0)
			})

			Convey("Then Hamming Distance should be under threshold ", func() {
				distance := HammingDistance(sun_big.PhashMatrix, sun_small.PhashMatrix)
				So(distance, ShouldBeLessThanOrEqualTo, Treshold)
				distance = HammingDistance(sun_big.PhashMatrix, sun_small.PhashMatrix)
				So(distance, ShouldBeLessThanOrEqualTo, Treshold)
			})

		})

		Convey("When CDCT Hashes are computed", func() {
			sun_big.ComputeImageHashPhash(false)
			sun_small.ComputeImageHashPhash(false)

			Convey("Then the hashes should not be zero", func() {
				So(sun_big.CPhash, ShouldNotEqual, 0)
				So(sun_small.CPhash, ShouldNotEqual, 0)
			})

			Convey("Then Hamming Distance should be under threshold ", func() {
				distance := HammingDistance(sun_big.CPhash, sun_small.CPhash)
				So(distance, ShouldBeLessThanOrEqualTo, Treshold)
				distance = HammingDistance(sun_big.CPhash, sun_small.CPhash)
				So(distance, ShouldBeLessThanOrEqualTo, Treshold)
			})

		})

		Convey("When Radon Hashes are computed", func() {
			sun_big.ComputeImageHashRadon(false)
			sun_small.ComputeImageHashRadon(false)

			Convey("Then the hashes should not be zero", func() {
				So(len(sun_big.Digest.RadonDigest.Coeffs), ShouldNotEqual, 0)
				So(sun_big.Digest.RadonDigest.Coeffs, ShouldNotBeNil)
				So(len(sun_small.Digest.RadonDigest.Coeffs), ShouldNotEqual, 0)
				So(sun_small.Digest.RadonDigest.Coeffs, ShouldNotBeNil)
			})

			SkipConvey("Then the Cross Correlation for threshold should be true", func() {

				So(CrossCorr(sun_big.Digest.RadonDigest, sun_small.Digest.RadonDigest, -0.1), ShouldBeTrue)
				So(CrossCorr(sun_big.Digest.RadonDigest, sun_small.Digest.RadonDigest, 0.0), ShouldBeTrue)
			})

		})

	})

}

func TestDifferentImagesDoNotMatch(t *testing.T) {

	Convey("Given some cats, suns and girls are loaded", t, func() {

		cat_big := getImgBag(cats_dir, "cat_big.jpg", 0)
		cat_medium := getImgBag(cats_dir, "cat_medium.jpg", 0)
		cat_small := getImgBag(cats_dir, "cat_small.jpg", 0)
		cat_sky := getImgBag(cats_dir, "cat_sky.jpg", 0)
		cat_smile := getImgBag(cats_dir, "smiling.jpg", 0)
		cats := []*ImageBag{cat_big, cat_medium, cat_small, cat_sky, cat_smile}

		lena_toshop := getImgBag(lena_dir, "lena_std.tiff", 0)
		lena_nkd := getImgBag(lena_dir, "l_hires.jpg", 0)
		lena_vintage := getImgBag(lena_dir, "lena.bmp", 0)
		girls := []*ImageBag{lena_toshop, lena_nkd, lena_vintage}

		sun_big := getImgBag(cats_dir, "sun_big.jpg", 0)
		sun_small := getImgBag(cats_dir, "sun_small.jpg", 0)
		suns := []*ImageBag{sun_big, sun_small}

		Convey("When DCT Hashes are computed", func() {

			cat_big.ComputeGreyscaleDct(false)
			cat_medium.ComputeGreyscaleDct(false)
			cat_small.ComputeGreyscaleDct(false)
			cat_sky.ComputeGreyscaleDct(false)
			cat_smile.ComputeGreyscaleDct(false)

			lena_toshop.ComputeGreyscaleDct(false)
			lena_nkd.ComputeGreyscaleDct(false)
			lena_vintage.ComputeGreyscaleDct(false)

			sun_big.ComputeGreyscaleDct(false)
			sun_small.ComputeGreyscaleDct(false)

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

			cat_big.ComputeGreyscaleDctMatrix(false)
			cat_medium.ComputeGreyscaleDctMatrix(false)
			cat_small.ComputeGreyscaleDctMatrix(false)
			cat_sky.ComputeGreyscaleDctMatrix(false)
			cat_smile.ComputeGreyscaleDctMatrix(false)

			lena_toshop.ComputeGreyscaleDctMatrix(false)
			lena_nkd.ComputeGreyscaleDctMatrix(false)
			lena_vintage.ComputeGreyscaleDctMatrix(false)

			sun_big.ComputeGreyscaleDctMatrix(false)
			sun_small.ComputeGreyscaleDctMatrix(false)

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

			cat_big.ComputeImageHashPhash(false)
			cat_medium.ComputeImageHashPhash(false)
			cat_small.ComputeImageHashPhash(false)
			cat_sky.ComputeImageHashPhash(false)
			cat_smile.ComputeImageHashPhash(false)

			// lena_toshop.ComputeImageHashPhash(false)
			lena_vintage.ComputeImageHashPhash(false)
			lena_nkd.ComputeImageHashPhash(false)
			girls := []*ImageBag{lena_vintage, lena_nkd}

			sun_big.ComputeImageHashPhash(false)
			sun_small.ComputeImageHashPhash(false)

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

			cat_big.ComputeImageHashRadon(false)
			cat_medium.ComputeImageHashRadon(false)
			cat_small.ComputeImageHashRadon(false)
			cat_sky.ComputeImageHashRadon(false)
			cat_smile.ComputeImageHashRadon(false)

			lena_toshop.ComputeImageHashRadon(false)
			lena_nkd.ComputeImageHashRadon(false)
			lena_vintage.ComputeImageHashRadon(false)

			sun_big.ComputeImageHashRadon(false)
			sun_small.ComputeImageHashRadon(false)

			Convey("Then cats do not look like girls", func() {
				PairExecuteFor2ImagesList(cats, girls, func(cat, girl *ImageBag) {
					So(CrossCorr(cat.Digest.RadonDigest, girl.Digest.RadonDigest, 0.8), ShouldBeFalse)
				})
			})

			Convey("And suns do not look like cats", func() {
				PairExecuteFor2ImagesList(suns, cats, func(sun, cat *ImageBag) {
					So(CrossCorr(sun.Digest.RadonDigest, cat.Digest.RadonDigest, 0.8), ShouldBeFalse)
				})
			})

			Convey("And girls do not look like suns", func() {
				PairExecuteFor2ImagesList(girls, suns, func(girl, sun *ImageBag) {
					So(CrossCorr(girl.Digest.RadonDigest, sun.Digest.RadonDigest, 0.8), ShouldBeFalse)
				})
			})

		})

	})

}
