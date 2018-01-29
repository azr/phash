package phash

import (
	"image"
	"image/color"

	"github.com/azr/phash/cmd"
	"github.com/azr/phash/cornerdetect"

	"github.com/azr/gift"
)

const (
	kParam = 0.04
)

// Points is a slice of points.
type Points []image.Point

func (p Points) Len() int           { return len(p) }
func (p Points) Less(i, j int) bool { return p[i].X < p[j].X }
func (p Points) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

func detectCorners(src *image.Gray) []image.Point {
	pts := cornerdetect.Fast9(src, 30)
	pixels := len(src.Pix)
	end := min(pixels/400, len(pts)-1)
	return pts[:end]
}

// FindKeypoints returns a list of points that are key points
func FindKeypoints(src image.Image) Points {
	g := gift.New(
		gift.Grayscale(),
	)

	simplifiedImage := image.NewGray(g.Bounds(src.Bounds()))
	g.Draw(simplifiedImage, src)
	points := detectCorners(simplifiedImage)
	for _, point := range points {
		simplifiedImage.Set(point.X+2, point.Y+2, color.Gray{255})
		simplifiedImage.Set(point.X+1, point.Y+1, color.Gray{255})
		simplifiedImage.Set(point.X-1, point.Y-1, color.Gray{255})
		simplifiedImage.Set(point.X-2, point.Y-2, color.Gray{255})
		simplifiedImage.Set(point.X+1, point.Y-1, color.Gray{255})
		simplifiedImage.Set(point.X-1, point.Y+1, color.Gray{255})
		simplifiedImage.Set(point.X, point.Y, color.Gray{255})
	}
	cmd.WriteImageToPath(simplifiedImage, "./points")
	return points
}
