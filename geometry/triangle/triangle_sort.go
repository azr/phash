package triangle

import "sort"

func (t *Triangle) Len() int      { return 3 }
func (t *Triangle) Swap(x, y int) { t[x], t[y] = t[y], t[x] }

func (t *Triangle) Less(x, y int) bool {
	if t[x].X == t[y].X {
		return t[x].Y < t[y].Y
	}
	return t[x].X < t[y].X
}

var _ sort.Interface = &Triangle{}

// Triangles is a sortable array of triangles
type Triangles []Triangle

func (ts Triangles) Len() int           { return len(ts) }
func (ts Triangles) Swap(x, y int)      { ts[x], ts[y] = ts[y], ts[x] }
func (ts Triangles) Diff(x, y int) bool { return ts[x] != ts[y] }

func (ts Triangles) Less(x, y int) bool {
	for i := 0; i < 3; i++ {
		if ts[x][i].X == ts[y][i].X &&
			ts[x][i].Y < ts[y][i].Y {
			return true
		}
		if ts[x][i].X < ts[y][i].X {
			return true
		}
	}
	return false
}

var _ sort.Interface = &Triangles{}

// Unique sorts input and and returns it
// with duplicates removed.
// Unique expects all triangles to have points sorted out.
// it will sort
func Unique(input Triangles) []Triangle {
	if input.Len() <= 1 {
		return input
	}
	sort.Sort(input)
	j := 1
	i := 0
	for ; j < input.Len(); j++ {
		if input.Diff(i, j) {
			i++
			input[i] = input[j]
		}
	}
	if i != j-1 && input[i] == input[j-1] {
		j--
	}
	return input[:j]
}
