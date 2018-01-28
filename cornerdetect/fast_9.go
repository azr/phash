package cornerdetect

import (
	"image"
	"sort"
)

// Fast9 corner detection gives you sorted corners using
// transpiled code from C https://github.com/edrosten/fast-C-src
//
// The higher the threshold the mor picky it gets.
// Points are sorted by score higher scores at the top.
func Fast9(src *image.Gray, treshold int) []image.Point {
	bounds := src.Bounds()
	return fast9_detect_nonmax(src.Pix, bounds.Max.X, bounds.Max.Y, src.Stride, treshold)
}

type cornerScoreSorter struct {
	corners []image.Point
	scores  []int
}

func (css cornerScoreSorter) Len() int { return len(css.scores) }
func (css cornerScoreSorter) Swap(i, j int) {
	css.scores[i], css.scores[j] = css.scores[j], css.scores[i]
	css.corners[i], css.corners[j] = css.corners[j], css.corners[i]
}
func (css cornerScoreSorter) Less(i, j int) bool { return css.scores[i] > css.scores[j] }

func fast9_detect_nonmax(im []uint8, xsize int, ysize int, stride int, threshold int) []image.Point {
	corners := fast9_detect(im, xsize, ysize, stride, threshold)
	scores := fast9_score(im, stride, corners, threshold)
	nonmax := nonmax_suppression(corners, scores)
	sort.Sort(nonmax)
	return nonmax.corners
}
