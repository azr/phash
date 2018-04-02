package phash

import (
	"image"
	"log"
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

func (a Matrix) at(x int) float64 {
	if a.determinant == 0 {
		a.determinant = 1
	}
	return a.m[x] / a.determinant
}

func (a Matrix) Mul(b Matrix) Matrix {
	return Matrix{
		m: [9]float64{
			a.at(0)*b.at(0) + a.at(1)*b.at(3) + a.at(2)*b.at(6), a.at(0)*b.at(1) + a.at(1)*b.at(4) + a.at(2)*b.at(7), a.at(0)*b.at(2) + a.at(1)*b.at(5) + a.at(2)*b.at(8),
			a.at(3)*b.at(0) + a.at(4)*b.at(3) + a.at(5)*b.at(6), a.at(3)*b.at(1) + a.at(4)*b.at(4) + a.at(5)*b.at(7), a.at(3)*b.at(2) + a.at(4)*b.at(5) + a.at(5)*b.at(8),
			a.at(6)*b.at(0) + a.at(7)*b.at(3) + a.at(8)*b.at(6), a.at(6)*b.at(1) + a.at(7)*b.at(4) + a.at(8)*b.at(7), a.at(6)*b.at(2) + a.at(7)*b.at(5) + a.at(8)*b.at(8),
		},
	}
}

func (a Matrix) Mul1(b [3]float64) (x, y, z float64) {
	return a.at(0)*b[0] + a.at(1)*b[1] + a.at(2)*b[2],
		a.at(3)*b[0] + a.at(4)*b[1] + a.at(5)*b[2],
		a.at(6)*b[0] + a.at(7)*b[1] + a.at(8)*b[2]
}

func (a Matrix) TransformPoint(x, y int) (int, int) {
	xp, yp, _ := a.Mul1([3]float64{float64(x), float64(y), 1})
	return int(xp), int(yp)
}

func (a Triangle) ExtractEquilateralTriangleFrom(src image.Image) image.Image {
	res := image.NewRGBA(image.Rect(0, 0, eTriangleSize, eTriangleSize))

	invB := a.InverseMatrix()
	b := eTriangleMatrix
	m := b.Mul(invB)

	for x := 0; x < eTriangleSize; x++ {
		for y := 0; y < eTriangleHeight; y++ {
			if !eTriangle.Contains(x, y) {
				continue
			}
			// for each point of eTriangle
			// find correspondig pixel in triangle a &
			xa, ya := m.TransformPoint(x, y)
			// set it in triangle b
			res.Set(x, y, src.At(xa, ya))
			log.Printf("set %d, %d, from %d, %d", x, y, xa, ya)
		}
	}
	return res
}

func (a Triangle) Determinant() int {
	return a.B().X*a.C().Y - a.C().X*a.B().Y - a.A().Y*a.B().X + a.A().Y*a.C().X + a.A().X*a.B().Y - a.A().X*a.C().Y
}

func (a Triangle) InverseMatrix() Matrix { // TODO: check me because I might be wrong
	cofactorMatrix := [9]float64{
		float64(a.B().Y - a.C().Y), -float64(a.A().Y - a.C().Y), float64(a.A().Y - a.B().Y),
		-float64(a.B().X - a.C().X), float64(a.A().X - a.C().X), -float64(a.A().X - a.B().X),
		float64((a.B().Y * a.C().Y) - (a.B().Y * a.C().X)), -float64((a.A().X * a.C().Y) - (a.A().Y * a.C().X)), float64((a.A().X * a.B().Y) - (a.A().Y * a.B().X)),
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
