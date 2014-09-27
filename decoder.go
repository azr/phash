package phash

import (
	"fmt"
	"github.com/nfnt/resize"
	"image"
	// Package image/[jpeg|fig|png] is not used explicitly in the code below,
	// but is imported for its initialization side-effect, which allows
	// image.Decode to understand [jpeg|gif|png] formatted images.
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"io"
)

//Digest will contain digested dct/radon/other digest
//of an image
type Digest struct {
	Image       image.Image
	Format      string
	Phash       uint64
	PhashMatrix uint64
	RadonDigest RadonDigest
}

//ComputeGreyscaleDct puts the result of GreyscaleDct in a digest
func (d *Digest) ComputeGreyscaleDct() error {
	stamp := resize.Resize(32, 32, d.Image, resize.Bilinear)
	greyscaleStamp := Gscl(stamp)

	// greyscaleStamp := greyscale.Greyscale(stamp)
	d.Phash = GreyscaleDct(greyscaleStamp)

	return nil
}

//ComputeGreyscaleDctMatrix puts the result of GreyscaleDctMatrix in a digest
func (d *Digest) ComputeGreyscaleDctMatrix() error {
	stamp := resize.Resize(32, 32, d.Image, resize.Bilinear)
	greyscaleStamp := Gscl(stamp)

	// greyscaleStamp := greyscale.Greyscale(stamp)
	d.PhashMatrix = GreyscaleDctMatrix(greyscaleStamp)

	return nil
}

//ComputeRadonDigest puts the result of Radon in a digest
func (d *Digest) ComputeRadonDigest() error {
	imgMtx, err := imageToMatrix(d.Image)
	if err != nil {
		panic(err)
	}

	d.RadonDigest = Radon(imgMtx)

	return nil
}

//DecodeAndDigest reads an image and fills all the fields of a Digest
func DecodeAndDigest(r io.Reader) (digest *Digest, err error) {
	digest = new(Digest)
	if digest == nil {
		return nil, fmt.Errorf("Could not allocate Digest memory")
	}

	if digest.Image, digest.Format, err = image.Decode(r); err != nil {
		return
	}

	digest.ComputeGreyscaleDctMatrix()
	digest.ComputeGreyscaleDct()
	digest.ComputeRadonDigest()

	return
}
