package euler

import (
	"fmt"
	"strings"
	"testing"
)

// Used in p18 and p67
func P18ParseTriangle(tri string) [][]int {

	lines := strings.Split(tri, "\n")

	data := make([][]int, 0, len(lines))
	for idxl, l := range lines {
		if len(l) == 0 {
			continue
		}
		nums := strings.Split(l, " ")
		dataline := make([]int, len(nums))
		for idx, n := range nums {
			dataline[idx] = int(n[0]-'0')*10 + int(n[1]-'0')
		}
		data = append(data, dataline)

		if idxl+1 != len(dataline) {
			panic(fmt.Errorf("error parsing triangle: %d, %d", idxl+1, len(dataline)))
		}
	}
	return data
}

func P18MaxPathSum(tri [][]int) int {

	line := len(tri) - 1
	best := make([]int, line+1)
	copy(best, tri[line])

	for line = line - 1; line >= 0; line-- {
		newbest := make([]int, line+1)
		for i := 0; i < len(newbest); i++ {
			max := best[i]
			if max < best[i+1] {
				max = best[i+1]
			}
			newbest[i] = tri[line][i] + max
		}
		best = newbest
	}
	return best[0]
}

func TestP18(t *testing.T) {
	data := P18ParseTriangle(p18input)
	total := P18MaxPathSum(data)
	if total != 1074 {
		t.Error(total)
	}
}

const p18input = `75
95 64
17 47 82
18 35 87 10
20 04 82 47 65
19 01 23 75 03 34
88 02 77 73 07 63 67
99 65 04 28 06 16 70 92
41 41 26 56 83 40 80 70 33
41 48 72 33 47 32 37 16 94 29
53 71 44 65 25 43 91 52 97 51 14
70 11 33 28 77 73 17 78 39 68 17 57
91 71 52 38 17 14 91 43 58 50 27 29 48
63 66 04 68 89 53 67 30 73 16 69 87 40 31
04 62 98 27 23 09 70 98 73 93 38 53 60 04 23
`
