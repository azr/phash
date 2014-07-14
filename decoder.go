
package phash

import (
	"io"
	"fmt"
	"image"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"github.com/nfnt/resize"
)

type Digest struct {
	Image		image.Image
	Format		string
	Phash		uint64
	PhashMatrix	uint64
	RadonDigest	RadonDigest
}

func (d *Digest) ComputeGreyscaleDct() {
	stamp := resize.Resize(32, 32, d.Image, resize.Bilinear)
	greyscaleStamp := Gscl(stamp)

	// greyscaleStamp := greyscale.Greyscale(stamp)
	d.Phash = GreyscaleDct(greyscaleStamp)
}

func (d *Digest) ComputeGreyscaleDctMatrix() {
	stamp := resize.Resize(32, 32, d.Image, resize.Bilinear)
	greyscaleStamp := Gscl(stamp)

	// greyscaleStamp := greyscale.Greyscale(stamp)
	d.PhashMatrix = GreyscaleDctMatrix(greyscaleStamp)
}

func Decode(r io.Reader) (digest *Digest, err error) {
	digest = new(Digest)
	if digest == nil {
		return nil, fmt.Errorf("Could not allocate Digest memory")
	}

	if digest.Image, digest.Format, err = image.Decode(r); err != nil {
		return
	}

	return
}
