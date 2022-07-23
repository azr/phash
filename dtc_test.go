package phash

import (
	_ "embed"
	"image"
	"strings"
	"testing"
)

//go:embed testdata/lena/l_hires.jpg
var lHiresJPGFile string

var (
	lHiresJPG, _, _ = image.Decode(strings.NewReader(lHiresJPGFile))
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
		// This test is quite handy, because I could find the expected
		// result in the research papers. A picture a bit more SFW
		// would have been better though.
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
