package phash

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

func (a *Matrix) Mul1(b [3]float64) (x, y, z float64) {
	a.fixDet()
	return a.m[0]*b[0] + a.m[1]*b[1] + a.m[2]*b[2],
		a.m[3]*b[0] + a.m[4]*b[1] + a.m[5]*b[2],
		a.m[6]*b[0] + a.m[7]*b[1] + a.m[8]*b[2]
}

func (a *Matrix) TransformPoint(x, y int) (int, int) {
	xp, yp, _ := a.Mul1([3]float64{float64(x), float64(y), 1})
	return int(xp), int(yp)
}

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

type equilateralTriangleExtract struct {
	src          image.Image
	fromTriangle Triangle
	m            *Matrix
}

func (_ equilateralTriangleExtract) Bounds() image.Rectangle {
	return image.Rect(0, 0, eTriangleSize, eTriangleSize)
}

func (e equilateralTriangleExtract) ColorModel() color.Model { return e.src.ColorModel() }
func (e equilateralTriangleExtract) At(x, y int) color.Color {
	if !eTriangle.Contains(x, y) {
		return color.Black
	}
	return e.src.At(e.m.TransformPoint(x, y))
}

func (a Triangle) Determinant() int {
	return a.B().X*a.C().Y - a.C().X*a.B().Y - a.A().Y*a.B().X + a.A().Y*a.C().X + a.A().X*a.B().Y - a.A().X*a.C().Y
}

func (a Triangle) InverseMatrix() Matrix {
	cofactorMatrix := [9]float64{
		float64(a.B().Y - a.C().Y), -float64(a.B().X - a.C().X), float64((a.B().X * a.C().Y) - (a.B().Y * a.C().X)),
		-float64(a.A().Y - a.C().Y), float64(a.A().X - a.C().X), -float64((a.A().X * a.C().Y) - (a.A().Y * a.C().X)),
		float64(a.A().Y - a.B().Y), -float64(a.A().X - a.B().X), float64((a.A().X * a.B().Y) - (a.A().Y * a.B().X)),
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
