package manipulator

import (
	"image/color"
	"strconv"
	"testing"
)

func TestColor(t *testing.T) {
	baseColor := color.RGBA{251, 252, 253, 254}
	hex := ColorToHexString(baseColor)
	n, _ := strconv.ParseInt(hex, 0, 64)

	f64 := ColorToFloat64(baseColor)

	if n != int64(f64) {
		t.Fatal("Color convertion not equal ?")
	}

	recolor := Float64ToColor(f64)
	r, g, b, _ := baseColor.RGBA()
	rr, rg, rb, _ := recolor.RGBA()
	if rr != r || rg != g || rb != b {
		t.Fatal("Cant convert back to color !, ", rr, r, rg, g, rb, b)
	}
}
