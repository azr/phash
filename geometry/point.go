package geometry

import (
	"image"

	"github.com/azr/phash/cornerdetect"
)

// Points is a slice of points.
type Points []image.Point

func (p Points) Len() int           { return len(p) }
func (p Points) Less(i, j int) bool { return p[i].X < p[j].X }
func (p Points) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

// DetectCorners gives you a list of points containing
// corner of the image src
// src needs to be a grayscale image
func DetectCorners(src *image.Gray) []image.Point {
	pts := cornerdetect.Fast9(src, 30)
	pixels := len(src.Pix)
	end := Min(pixels/400, len(pts)-1)
	return pts[:end]
}
