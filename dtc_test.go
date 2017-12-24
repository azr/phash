package phash

import (
	"image"
	"testing"

	"github.com/azr/phash/cmd"
)

var (
	lHiresJPG, _ = cmd.OpenImageFromPath("/Users/azr/go/src/github.com/azr/phash/testdata/lena/l_hires.jpg")
	lHiresBMP, _ = cmd.OpenImageFromPath("/Users/azr/go/src/github.com/azr/phash/testdata/lena/lena.bmp")
)

func TestDTC(t *testing.T) {
	type args struct {
		img image.Image
	}
	tests := []struct {
		name      string
		args      args
		wantPhash uint64
	}{
		{name: "lena jpg",
			args:      args{img: lHiresJPG},
			wantPhash: parseBinary("0100011111110101011000000111111011000110010001110001010100011000")},
		{name: "zero",
			args:      args{img: nil},
			wantPhash: 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotPhash := DTC(tt.args.img); gotPhash != tt.wantPhash {
				t.Errorf("DTC() =\n%064b\nwant :\n%064b", gotPhash, tt.wantPhash)
			}
		})
	}
}
