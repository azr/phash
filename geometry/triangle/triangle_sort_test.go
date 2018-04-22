package triangle

import (
	"image"
	"reflect"
	"testing"
)

func TestUnique(t *testing.T) {
	type args struct {
		input Triangles
	}
	tests := []struct {
		name string
		args args
		want []Triangle
	}{
		{name: "len 1", args: args{
			input: Triangles{
				{image.Point{1, 2}},
			}},
			want: Triangles{
				{image.Point{1, 2}},
			},
		},
		{name: "len 2, sorted, no dupes", args: args{
			input: Triangles{
				{image.Point{1, 2}},
				{image.Point{3, 4}},
			}},
			want: Triangles{
				{image.Point{1, 2}},
				{image.Point{3, 4}},
			},
		},
		{name: "len 2, unsorted, no dupes", args: args{
			input: Triangles{
				{image.Point{3, 4}},
				{image.Point{1, 2}},
			}},
			want: Triangles{
				{image.Point{1, 2}},
				{image.Point{3, 4}},
			},
		},
		{name: "len 2,  dupes", args: args{
			input: Triangles{
				{image.Point{1, 2}},
				{image.Point{1, 2}},
			}},
			want: Triangles{
				{image.Point{1, 2}},
			},
		},
		{name: "len 3, unsorted, no dupes", args: args{
			input: Triangles{
				{image.Point{3, 2}},
				{image.Point{2, 2}},
				{image.Point{1, 2}},
			}},
			want: Triangles{
				{image.Point{1, 2}},
				{image.Point{2, 2}},
				{image.Point{3, 2}},
			},
		},
		{name: "len 3, unsorted, dupes", args: args{
			input: Triangles{
				{image.Point{2, 2}},
				{image.Point{2, 2}},
				{image.Point{1, 2}},
			}},
			want: Triangles{
				{image.Point{1, 2}},
				{image.Point{2, 2}},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := Unique(tt.args.input)
			if !reflect.DeepEqual(res, tt.want) {
				t.Errorf("Unique() = %v, want %v", res, tt.want)
			}
		})
	}
}
