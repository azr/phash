package geometry

import (
	"testing"
)

func Test_image2DTransformInt_mul(t *testing.T) {
	type fields struct {
		m [4]int
	}
	type args struct {
		x int
		y int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		wantX  int
		wantY  int
	}{
		{name: "first", fields: fields{m: rota90M}, args: args{1, 0}, wantX: 0, wantY: 1},
		{name: "basic", fields: fields{m: rota90M}, args: args{3, 4}, wantX: -4, wantY: 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i2dt := &image2DTransformInt{
				m: tt.fields.m,
			}
			got, got1 := i2dt.m.mul(tt.args.x, tt.args.y)
			if got != tt.wantX {
				t.Errorf("image2DTransformInt.mul() got = %v, want %v", got, tt.wantX)
			}
			if got1 != tt.wantY {
				t.Errorf("image2DTransformInt.mul() got1 = %v, want %v", got1, tt.wantY)
			}
		})
	}
}
