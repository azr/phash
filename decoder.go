package phash

import (
	"fmt"
	"github.com/nfnt/resize"
	"image"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"io"
)

type Digest struct {
	Image       image.Image
	Format      string
	Phash       uint64
	PhashMatrix uint64
	RadonDigest RadonDigest
}

func (d *Digest) ComputeGreyscaleDct() error {
	stamp := resize.Resize(32, 32, d.Image, resize.Bilinear)
	greyscaleStamp := Gscl(stamp)

	// greyscaleStamp := greyscale.Greyscale(stamp)
	d.Phash = GreyscaleDct(greyscaleStamp)

	return nil
}

func (d *Digest) ComputeGreyscaleDctMatrix() error {
	stamp := resize.Resize(32, 32, d.Image, resize.Bilinear)
	greyscaleStamp := Gscl(stamp)

	// greyscaleStamp := greyscale.Greyscale(stamp)
	d.PhashMatrix = GreyscaleDctMatrix(greyscaleStamp)

	return nil
}

func (d *Digest) ComputeGreyscaleRadonDigest() error {
	// stamp := resize.Resize(32, 32, d.Image, resize.Bilinear)
	greyscaleStamp := Gscl(d.Image)

	d.RadonDigest = Radon(greyscaleStamp)

	return nil
}

func Decode(r io.Reader) (digest *Digest, err error) {
	digest = new(Digest)
	if digest == nil {
		return nil, fmt.Errorf("Could not allocate Digest memory")
	}

	if digest.Image, digest.Format, err = image.Decode(r); err != nil {
		return
	}

	digest.ComputeGreyscaleDctMatrix()
	digest.ComputeGreyscaleDct()
	digest.ComputeGreyscaleRadonDigest()

	return
}
