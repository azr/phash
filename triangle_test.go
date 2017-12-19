package phash

import (
	"image"
	"reflect"
	"testing"
)

func TestPoints_EveryTriangles(t *testing.T) {
	type args struct {
		opts EveryTrianglesOpts
	}
	tests := []struct {
		name   string
		points Points
		args   args
		want   []Triangle
	}{

		{name: "2 lines of 5 point slice - minArea of 6 - some points",
			points: Points{
				image.Point{1, 2}, image.Point{2, 3}, image.Point{3, 4}, image.Point{4, 5}, image.Point{5, 6},
				image.Point{3, 1}, image.Point{4, 2}, image.Point{5, 3}, image.Point{6, 4}, image.Point{7, 5},
			}, args: args{opts: EveryTrianglesOpts{minArea: 6, upperThreshold: 999999}},
			want: []Triangle{
				Triangle{image.Point{1, 2}, image.Point{3, 1}, image.Point{7, 5}},
				Triangle{image.Point{2, 3}, image.Point{3, 1}, image.Point{7, 5}},
				Triangle{image.Point{3, 4}, image.Point{3, 1}, image.Point{7, 5}},
				Triangle{image.Point{4, 5}, image.Point{3, 1}, image.Point{7, 5}},
				Triangle{image.Point{5, 6}, image.Point{3, 1}, image.Point{7, 5}},
			},
		},

		// all following triangles are on a line, it makes testing easy for the mind...

		{name: "line of 5 point slice - minArea of 1 - no points", points: Points{image.Point{1, 2}, image.Point{2, 3}, image.Point{3, 4}, image.Point{4, 5}, image.Point{5, 6}}, args: args{opts: EveryTrianglesOpts{minArea: 1, upperThreshold: 999999}},
			want: []Triangle{},
		},

		{name: "line of 5 point slice - upperThreshold of 3", points: Points{image.Point{1, 2}, image.Point{2, 3}, image.Point{3, 4}, image.Point{4, 5}, image.Point{5, 6}}, args: args{opts: EveryTrianglesOpts{upperThreshold: 3}},
			want: []Triangle{
				Triangle{image.Point{1, 2}, image.Point{2, 3}, image.Point{3, 4}},
				Triangle{image.Point{2, 3}, image.Point{3, 4}, image.Point{4, 5}},
				Triangle{image.Point{3, 4}, image.Point{4, 5}, image.Point{5, 6}},
			},
		},
		{name: "line of 5 point slice - lowerThreshold of 2", points: Points{image.Point{1, 2}, image.Point{2, 3}, image.Point{3, 4}, image.Point{4, 5}, image.Point{5, 6}}, args: args{opts: EveryTrianglesOpts{lowerThreshold: 2, upperThreshold: 999999999}},
			want: []Triangle{
				Triangle{image.Point{1, 2}, image.Point{3, 4}, image.Point{5, 6}},
			},
		},
		{name: "line of 5 point slice", points: Points{image.Point{1, 2}, image.Point{2, 3}, image.Point{3, 4}, image.Point{4, 5}, image.Point{5, 6}}, args: args{opts: EveryTrianglesOpts{upperThreshold: 999999999}},
			want: []Triangle{
				Triangle{image.Point{1, 2}, image.Point{2, 3}, image.Point{3, 4}},
				Triangle{image.Point{1, 2}, image.Point{2, 3}, image.Point{4, 5}},
				Triangle{image.Point{1, 2}, image.Point{3, 4}, image.Point{4, 5}},
				Triangle{image.Point{1, 2}, image.Point{2, 3}, image.Point{5, 6}},
				Triangle{image.Point{1, 2}, image.Point{3, 4}, image.Point{5, 6}},
				Triangle{image.Point{1, 2}, image.Point{4, 5}, image.Point{5, 6}},
				Triangle{image.Point{2, 3}, image.Point{3, 4}, image.Point{4, 5}},
				Triangle{image.Point{2, 3}, image.Point{3, 4}, image.Point{5, 6}},
				Triangle{image.Point{2, 3}, image.Point{4, 5}, image.Point{5, 6}},
				Triangle{image.Point{3, 4}, image.Point{4, 5}, image.Point{5, 6}},
			},
		},
		{name: "line of 4 point slice", points: Points{image.Point{1, 2}, image.Point{2, 3}, image.Point{3, 4}, image.Point{4, 5}}, args: args{opts: EveryTrianglesOpts{upperThreshold: 999999999}},
			want: []Triangle{
				Triangle{image.Point{1, 2}, image.Point{2, 3}, image.Point{3, 4}},
				Triangle{image.Point{1, 2}, image.Point{2, 3}, image.Point{4, 5}},
				Triangle{image.Point{1, 2}, image.Point{3, 4}, image.Point{4, 5}},
				Triangle{image.Point{2, 3}, image.Point{3, 4}, image.Point{4, 5}},
			},
		},
		{name: "line of 3 point slice", points: Points{image.Point{1, 2}, image.Point{2, 3}, image.Point{3, 4}}, args: args{opts: EveryTrianglesOpts{upperThreshold: 999999999}},
			want: []Triangle{
				Triangle{image.Point{1, 2}, image.Point{2, 3}, image.Point{3, 4}},
			},
		},
		{name: "nil", points: Points{}, args: args{opts: EveryTrianglesOpts{}},
			want: []Triangle{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.points.EveryTriangles(tt.args.opts); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Points.TrianglesToFirst() = %v, want %v", got, tt.want)
			}
		})
	}
}
