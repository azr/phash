package triangle

import (
	"image"
	"image/color"
	"math"

	"github.com/disintegration/imaging"

	"github.com/StephaneBunel/bresenham"
	"github.com/azr/gift"
	"github.com/azr/phash/cmd"
)

// A Triangle is made of three 2D points.
// A triangle is also a 3x3 matrix:
// | x1, x2, x3 |
// | y1, y2, y3 |
// |  1,  1,  1 |
type Triangle [3]image.Point

func (t *Triangle) A() image.Point { return t[0] }
func (t *Triangle) B() image.Point { return t[1] }
func (t *Triangle) C() image.Point { return t[2] }

// Area of the triangle
func (t *Triangle) Area() int {
	x0 := t.A().X
	x1 := t.B().X
	x2 := t.C().X

	y0 := t.A().Y
	y1 := t.B().Y
	y2 := t.C().Y

	return (x0*(y1-y2) + x1*(y2-y0) + x2*(y0-y1)) / 2
}

// Debug outputs your image to a triangle.jpg image
// containing only triangle coordinates.
func (t *Triangle) Debug(img image.Image) {
	g := gift.New(
		gift.Mean(5, true),
	)

	// STUFF
	simplifiedImage := image.NewRGBA(g.Bounds(img.Bounds()))
	g.Draw(simplifiedImage, img)
	for _, p := range t {
		simplifiedImage.Set(p.X, p.Y, color.RGBA{B: 255, R: 255})
		simplifiedImage.Set(p.X+1, p.Y+1, color.RGBA{B: 255, R: 255})
		simplifiedImage.Set(p.X-1, p.Y-1, color.RGBA{B: 255, R: 255})
		simplifiedImage.Set(p.X-1, p.Y+1, color.RGBA{B: 255, R: 255})
		simplifiedImage.Set(p.X+1, p.Y-1, color.RGBA{B: 255, R: 255})
	}
	bresenham.Bresenham(simplifiedImage, t[0].X, t[0].Y, t[1].X, t[1].Y, color.RGBA{R: 255}) // draw line
	bresenham.Bresenham(simplifiedImage, t[2].X, t[2].Y, t[1].X, t[1].Y, color.RGBA{R: 255})
	bresenham.Bresenham(simplifiedImage, t[2].X, t[2].Y, t[0].X, t[0].Y, color.RGBA{R: 255})

	// make triangle
	tbounds := t.Bounds()
	triangle := image.NewRGBA(image.Rect(0, 0, tbounds.Dx(), tbounds.Dy()))
	for x := 0; x < tbounds.Max.X; x++ {
		for y := 0; y < tbounds.Max.Y; y++ {
			if !t.Contains(x, y) {
				continue
			}
			triangle.Set(x-(tbounds.Min.X), y-(tbounds.Min.Y), simplifiedImage.At(x, y))
			// simplifiedImage.Set(x, y, color.Black)
		}
	}
	cmd.WriteImageToPath(simplifiedImage, "full-triangle")
	cmd.WriteImageToPath(triangle, "triangle")
	subImg := simplifiedImage.SubImage(tbounds)
	cmd.WriteImageToPath(subImg, "square-triangle")
	angle := math.Atan2(float64(t[0].X-t[1].X), float64(t[0].Y-t[1].Y))
	angle = math.Abs(angle)

	rotatedTriangle := imaging.Rotate(triangle, angle*180/math.Pi, color.Black)
	rotated := imaging.Rotate(subImg, angle*180/math.Pi, color.Black)

	cmd.WriteImageToPath(rotated, "rot")
	cmd.WriteImageToPath(rotatedTriangle, "rot-tri")
	println("angle :", angle)
	println("angle :", angle)
}

func sign(p1, p2, p3 image.Point) bool {
	return (p1.X-p3.X)*(p2.Y-p3.Y)-(p2.X-p3.X)*(p1.Y-p3.Y) < 0
}

// Contains point x,y ?
func (t *Triangle) Contains(x, y int) bool {
	pt := image.Point{x, y}
	b1 := sign(pt, t[0], t[1])
	b2 := sign(pt, t[1], t[2])
	b3 := sign(pt, t[2], t[0])
	return ((b1 == b2) && (b2 == b3))
}

const triangleScale = 160

// FlatVerticeImage returns a triangular image that references img
// points that are not in triangle will be removed
// vertice can be 1, 2 or 3
func (t *Triangle) FlatVerticeImage(img image.Image, vertice int) image.Image {

	bounds := t.Bounds()
	subImg := img.(interface {
		SubImage(r image.Rectangle) image.Image
	}).SubImage(bounds)

	// I
	// 1. get copy of resized image
	// 2. set outside pixels to black
	// II
	// 1. use a reference to top image
	// points will reference image in a resize
	// way.
	return subImg
}

// Bounds returns a rectangle containing triangle
func (t *Triangle) Bounds() (res image.Rectangle) {
	res.Min.Y, res.Min.X = t[0].Y, t[0].X
	res.Max.Y, res.Max.X = t[0].Y, t[0].X

	for i := 1; i < 3; i++ {
		point := t[i]
		if point.X < res.Min.X {
			res.Min.X = point.X
		}
		if point.X > res.Max.X {
			res.Max.X = point.X
		}
		if point.Y < res.Min.Y {
			res.Min.Y = point.Y
		}
		if point.Y > res.Max.Y {
			res.Max.Y = point.Y
		}
	}
	// res.Min.Y -= 5
	// res.Min.X -= 5
	// res.Max.Y += 5
	// res.Max.X += 5
	return res
}
