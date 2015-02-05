package euler

import "testing"

func Min(args ...int64) int64 {
	min := args[0]
	for _, v := range args[1:] {
		if min > v {
			min = v
		}
	}
	return min
}

func TestP82(t *testing.T) {

	solve := func(matrix [][]int64, sz int) {

		mod := make([]int64, sz)
		for x := 1; x < sz; x++ {

			for y := 0; y < sz; y++ {
				mod[y] = matrix[y][x] + matrix[y][x-1]
			}

			changed := true
			for changed {
				changed = false
				for y := 0; y < sz; y++ {
					if y > 0 {
						top := mod[y-1] + matrix[y][x]
						if mod[y] > top {
							mod[y] = top
							changed = true
						}
					}
					if y < sz-1 {
						bottom := mod[y+1] + matrix[y][x]
						if mod[y] > bottom {
							mod[y] = bottom
							changed = true
						}
					}
				}
			}

			for y := 0; y < sz; y++ {
				matrix[y][x] = mod[y]
			}
		}
	}

	solution := func(matrix [][]int64, sz int) int64 {
		sol := matrix[0][sz-1]
		for y := 1; y < sz; y++ {
			if sol > matrix[y][sz-1] {
				sol = matrix[y][sz-1]
			}
		}
		return sol
	}

	matrix := p81TestMatrix()
	solve(matrix, 5)
	sol := solution(matrix, 5)
	if sol != 994 {
		t.Error(sol)
	}

	matrix = p81ParseMatrix(p81Matrix)
	solve(matrix, 80)
	sol = solution(matrix, 80)
	if sol != 260324 {
		t.Error(sol)
	}
}
