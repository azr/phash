package radon

import (
	"fmt"
	"github.com/azr/phash/manipulator"
	"image"
	// Package image/[jpeg|fig|png] is not used explicitly in the code below,
	// but is imported for its initialization side-effect, which allows
	// image.Decode to understand [jpeg|gif|png] formatted images.
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"io"
)

//ImageDigest will contain digested dct/radon/other digest
//of an image
type ImageDigest struct {
	Image  image.Image
	Format string
	Digest Digest
}

//ComputeRadonDigest puts the result of Radon in a digest
func (d *ImageDigest) ComputeRadonDigest() error {
	imgMtx, err := manipulator.ImageToMatrix(d.Image)
	if err != nil {
		panic(err)
	}

	d.Digest = DigestMatrix(imgMtx)

	return nil
}

//DecodeAndDigest reads an image and fills all the fields of a ImageDigest
func DecodeAndDigest(r io.Reader) (digest *ImageDigest, err error) {
	digest = new(ImageDigest)
	if digest == nil {
		return nil, fmt.Errorf("Could not allocate ImageDigest memory")
	}

	if digest.Image, digest.Format, err = image.Decode(r); err != nil {
		return
	}

	digest.ComputeRadonDigest()

	return
}
