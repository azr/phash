package triangle

import (
	"image"
)

// A Triangle is made of three 2D points.
// A triangle is also a 3x3 matrix:
// | x1, x2, x3 |
// | y1, y2, y3 |
// |  1,  1,  1 |
type Triangle [3]image.Point

// Area of the triangle
func (t *Triangle) Area() int {
	x0 := t[0].X
	x1 := t[1].X
	x2 := t[2].X

	y0 := t[0].Y
	y1 := t[1].Y
	y2 := t[2].Y

	return (x0*(y1-y2) + x1*(y2-y0) + x2*(y0-y1)) / 2
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
