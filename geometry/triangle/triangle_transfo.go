package triangle

import (
	"image"
	"image/color"
	"math"
)

const (
	eTriangleSize = 45
)

var (
	_eTriangleHeight = eTriangleSize * math.Sqrt(3) / 2
	eTriangleHeight  = int(_eTriangleHeight)
	eTriangle        = Triangle{
		image.Point{0, 0},
		image.Point{eTriangleSize / 2, eTriangleHeight},
		image.Point{eTriangleSize, 0},
	}
	eTriangleMatrix = Matrix{
		m: [9]float64{
			0, eTriangleSize / 2, eTriangleSize,
			0, _eTriangleHeight, 0,
			1, 1, 1,
		},
	}
)

// Matrix is a 3x3 matrix.
//
// A triangle in the matrix form is:
//      | aa.x ab.x ac.x |
//  A = | aa.y ab.y ac.y |
//      |  1    1    1   |
//
// To transform a triangle into another triangle
// wee need to find M from :
//  M * A = B
// which simplifies to :
//  M = B * Inv(A)
type Matrix struct {
	m           [9]float64
	determinant float64
}

// Mul return a * b which are 3/3 matrixes
func (a *Matrix) Mul(b *Matrix) Matrix {
	a.fixDet()
	b.fixDet()
	res := Matrix{
		m: [9]float64{
			a.m[0]*b.m[0] + a.m[1]*b.m[3] + a.m[2]*b.m[6], a.m[0]*b.m[1] + a.m[1]*b.m[4] + a.m[2]*b.m[7], a.m[0]*b.m[2] + a.m[1]*b.m[5] + a.m[2]*b.m[8],
			a.m[3]*b.m[0] + a.m[4]*b.m[3] + a.m[5]*b.m[6], a.m[3]*b.m[1] + a.m[4]*b.m[4] + a.m[5]*b.m[7], a.m[3]*b.m[2] + a.m[4]*b.m[5] + a.m[5]*b.m[8],
			a.m[6]*b.m[0] + a.m[7]*b.m[3] + a.m[8]*b.m[6], a.m[6]*b.m[1] + a.m[7]*b.m[4] + a.m[8]*b.m[7], a.m[6]*b.m[2] + a.m[7]*b.m[5] + a.m[8]*b.m[8],
		},
	}
	return res
}

func (a *Matrix) fixDet() {
	if a.determinant != 0 && a.determinant != 1 {
		for i := range a.m {
			a.m[i] = a.m[i] * (1 / a.determinant)
		}
		a.determinant = 1
	}
}

// Mul1 returns a ( 3 by 3 ) * b ( 1 by 3)
func (a *Matrix) Mul1(b [3]float64) (x, y, z float64) {
	a.fixDet()
	return a.m[0]*b[0] + a.m[1]*b[1] + a.m[2]*b[2],
		a.m[3]*b[0] + a.m[4]*b[1] + a.m[5]*b[2],
		a.m[6]*b[0] + a.m[7]*b[1] + a.m[8]*b[2]
}

// TransformPoint gives you where point will be after
// trasforming it through a
func (a *Matrix) TransformPoint(x, y int) (int, int) {
	xp, yp, _ := a.Mul1([3]float64{float64(x), float64(y), 1})
	return int(xp), int(yp)
}

// ExtractEquilateralTriangleFrom will give you an equilateral triangle
// extract from src at coordinates of a.
func (a Triangle) ExtractEquilateralTriangleFrom(src image.Image) image.Image {
	invA := a.InverseMatrix()
	b := eTriangleMatrix
	invA.fixDet()
	m := b.Mul(&invA)
	m.fixDet()

	return equilateralTriangleExtract{
		src:          src,
		fromTriangle: a,
		m:            &m,
	}
}

// equilateralTriangleExtract will display am equilateral
// triangle that references another triangle in an image
// When you call At extrapolated pixels will be used.
type equilateralTriangleExtract struct {
	src          image.Image
	fromTriangle Triangle
	m            *Matrix
}

func (equilateralTriangleExtract) Bounds() image.Rectangle {
	return image.Rect(0, 0, eTriangleSize, eTriangleSize)
}

func (e equilateralTriangleExtract) ColorModel() color.Model { return e.src.ColorModel() }
func (e equilateralTriangleExtract) At(x, y int) color.Color {
	if !eTriangle.Contains(x, y) {
		return color.Black
	}
	return e.src.At(e.m.TransformPoint(x, y))
}

// Determinant of triangle matrix
func (a Triangle) Determinant() int {
	return a[1].X*a[2].Y - a[2].X*a[1].Y - a[0].Y*a[1].X + a[0].Y*a[2].X + a[0].X*a[1].Y - a[0].X*a[2].Y
}

// InverseMatrix gives you the matrix inverse of a.
func (a Triangle) InverseMatrix() Matrix {
	cofactorMatrix := [9]float64{
		float64(a[1].Y - a[2].Y), -float64(a[1].X - a[2].X), float64((a[1].X * a[2].Y) - (a[1].Y * a[2].X)),
		-float64(a[0].Y - a[2].Y), float64(a[0].X - a[2].X), -float64((a[0].X * a[2].Y) - (a[0].Y * a[2].X)),
		float64(a[0].Y - a[1].Y), -float64(a[0].X - a[1].X), float64((a[0].X * a[1].Y) - (a[0].Y * a[1].X)),
	}
	xint := Matrix{
		m: [9]float64{
			cofactorMatrix[0], cofactorMatrix[1], cofactorMatrix[2],
			cofactorMatrix[3], cofactorMatrix[4], cofactorMatrix[5],
			cofactorMatrix[6], cofactorMatrix[7], cofactorMatrix[8],
		},
		determinant: float64(a.Determinant()),
	}
	return xint
}
