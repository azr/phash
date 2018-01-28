package cornerdetect

import "image"

// nonmax_suppression - transpiled function from  nonmax.c:7
// Point above points (roughly) to the pixel above the one of interest, if there
// is a feature there.
// Find where each row begins
//  (the corners are output in raster scan order). A beginning of -1 signifies
//  that there are no corners on that row.
// Check left
// Check right
// Check above (if there is a valid row above)
// Make sure that current point_above is one
//   row above.
// Make point_above point to the first of the pixels above the current point,
//   if it exists.
// Check below (if there is anything below)
// Nothing below
// Make point below point to one of the pixels belowthe current point, if it
//    exists.
func nonmax_suppression(corners []image.Point, scores []int) cornerScoreSorter {
	var last_row int
	var row_start []int
	var i int
	var j int
	var point_above int = 0
	var point_below int = 0
	num_corners := len(corners)
	if num_corners < 1 {
		return cornerScoreSorter{}
	}
	ret_nonmax := make([]image.Point, 0, num_corners)
	ret_scores := make([]int, 0, num_corners)
	last_row = corners[num_corners-1].Y
	row_start = make([]int, uint32(last_row+1))
	for i = 0; i < last_row+1; i++ {
		row_start[i] = -1
	}
	{
		var prev_row int = -1
		for i = 0; i < num_corners; i++ {
			if corners[i].Y != prev_row {
				row_start[corners[i].Y] = i
				prev_row = corners[i].Y
			}
		}
	}
	for i = 0; i < num_corners; i++ {
		var score int = scores[i]
		pos := corners[i]
		if i > 0 {
			if corners[i-1].X == pos.X-1 && corners[i-1].Y == pos.Y && scores[i-1] >= score {
				continue
			}
		}
		if i < num_corners-1 {
			if corners[i+1].X == pos.X+1 && corners[i+1].Y == pos.Y && scores[i+1] >= score {
				continue
			}
		}
		if pos.Y != 0 && row_start[pos.Y-1] != -1 {
			if corners[point_above].Y < pos.Y-1 {
				point_above = row_start[pos.Y-1]
			}
			for ; corners[point_above].Y < pos.Y && corners[point_above].X < pos.X-1; point_above++ {
			}
			for j = point_above; corners[j].Y < pos.Y && corners[j].X <= pos.X+1; j++ {
				x := corners[j].X
				if (x == pos.X-1 || x == pos.X || x == pos.X+1) && scores[j] >= score {
					goto cont
				}
			}
		}
		if pos.Y != last_row && row_start[pos.Y+1] != -1 && point_below < num_corners {
			if corners[point_below].Y < pos.Y+1 {
				point_below = row_start[pos.Y+1]
			}
			for ; point_below < num_corners && corners[point_below].Y == pos.Y+1 && corners[point_below].X < pos.X-1; point_below++ {
			}
			for j = point_below; j < num_corners && corners[j].Y == pos.Y+1 && corners[j].X <= pos.X+1; j++ {
				x := corners[j].X
				if (x == pos.X-1 || x == pos.X || x == pos.X+1) && scores[j] >= score {
					goto cont
				}
			}
		}
		ret_nonmax = append(ret_nonmax, corners[i])
		ret_scores = append(ret_scores, scores[i])
	cont:
	}
	return cornerScoreSorter{ret_nonmax, ret_scores}
}
