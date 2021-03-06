package triangle

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
			if got := tt.a.Mul(&tt.args.b); !reflect.DeepEqual(got, tt.want) {
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
			name: "basic test", a: NewTriangle(1, 1, 2, 2, 1, 3),
			want: Matrix{m: [9]float64{
				-1, -1, 4,
				2, 0, -2,
				-1, 1, 0}, determinant: 2},
		},
		{
			name: "real test", a: NewTriangle(3, 2272, 5, 2108, 84, 2184),
			want: Matrix{m: [9]float64{
				-76, 79, -166152,
				-88, -81, 184296,
				164, 2, -5036}, determinant: 13108},
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

func TestMatrix_Mul1(t *testing.T) {
	type fields struct {
		m           [9]float64
		determinant float64
	}
	type args struct {
		b [3]float64
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		wantX  float64
		wantY  float64
		wantZ  float64
	}{
		{name: "  0 * 0 = 0", fields: fields{m: [9]float64{0, 0, 0, 0, 0, 0, 0, 0, 0}}, args: args{[3]float64{0, 0, 0}}, wantX: 0, wantY: 0, wantZ: 0},
		{name: "  1 * 1 = 0", fields: fields{m: [9]float64{1, 1, 1, 1, 1, 1, 1, 1, 1}}, args: args{[3]float64{1, 1, 1}}, wantX: 3, wantY: 3, wantZ: 3},
		{name: "1-9 * 1 = 0", fields: fields{m: [9]float64{1, 2, 3, 4, 5, 6, 7, 8, 9}}, args: args{[3]float64{1, 1, 1}}, wantX: 6, wantY: 15, wantZ: 24},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := Matrix{
				m:           tt.fields.m,
				determinant: tt.fields.determinant,
			}
			gotX, gotY, gotZ := a.Mul1(tt.args.b)
			if gotX != tt.wantX {
				t.Errorf("Matrix.Mul1() gotX = %v, want %v", gotX, tt.wantX)
			}
			if gotY != tt.wantY {
				t.Errorf("Matrix.Mul1() gotY = %v, want %v", gotY, tt.wantY)
			}
			if gotZ != tt.wantZ {
				t.Errorf("Matrix.Mul1() gotZ = %v, want %v", gotZ, tt.wantZ)
			}
		})
	}
}
