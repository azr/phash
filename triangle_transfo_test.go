package phash

import (
	"image"
	"reflect"
	"testing"
)

func NewTriangle(x1, y1, x2, y2, x3, y3 int) Triangle {
	return Triangle{image.Point{x1, y1}, image.Point{x2, y2}, image.Point{x3, y3}}
}

func TestMatrix_Mul(t *testing.T) {
	type args struct {
		b Matrix
	}
	tests := []struct {
		name string
		a    Matrix
		args args
		want Matrix
	}{
		{
			name: "0 * 0 = 0", a: Matrix{m: [9]float64{0, 0, 0, 0, 0, 0, 0, 0, 0}}, args: args{Matrix{m: [9]float64{0, 0, 0, 0, 0, 0, 0, 0, 0}}},
			want: Matrix{m: [9]float64{0, 0, 0, 0, 0, 0, 0, 0, 0}},
		},
		{
			name: "1 * 1 = 3", a: Matrix{m: [9]float64{1, 1, 1, 1, 1, 1, 1, 1, 1}}, args: args{Matrix{m: [9]float64{1, 1, 1, 1, 1, 1, 1, 1, 1}}},
			want: Matrix{m: [9]float64{3, 3, 3, 3, 3, 3, 3, 3, 3}},
		},
		{
			name: "incremental numbers", a: Matrix{m: [9]float64{1, 2, 3, 4, 5, 6, 7, 8, 9}}, args: args{Matrix{m: [9]float64{10, 11, 12, 13, 14, 15, 16, 17, 18}}},
			want: Matrix{m: [9]float64{84, 90, 96, 201, 216, 231, 318, 342, 366}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.a.Mul(tt.args.b); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Matrix.Mul() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTriangle_Determinant(t *testing.T) {
	tests := []struct {
		name string
		a    Triangle
		want int
	}{
		{name: "zero", a: NewTriangle(0, 0, 0, 0, 0, 0), want: 0},
		{name: "6677", a: NewTriangle(1, 6, 44, 1, 55, 155), want: 6677},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.a.Determinant(); got != tt.want {
				t.Errorf("Triangle.Determinant() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTriangle_InverseMatrix(t *testing.T) {
	tests := []struct {
		name string
		a    Triangle
		want Matrix
	}{
		{
			name: "basic test :D", a: NewTriangle(1, 1, 2, 2, 1, 3),
			want: Matrix{m: [9]float64{-1, 2, -1, -1, 0, 1, 4, -2, 0}, determinant: 2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.a.InverseMatrix(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Triangle.InverseMatrix() = %v, want %v", got, tt.want)
			}
		})
	}
}
