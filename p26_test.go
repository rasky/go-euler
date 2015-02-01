package euler

import (
	"reflect"
	"testing"
)

func p26LenPeriod(x int) int {

	seen := make(map[int]int)
	num := 1
	count := 0
	for num != 0 {
		if idx, found := seen[num]; found {
			return count - idx
		}
		seen[num] = count

		if num < x {
			num *= 10
			continue
		}

		num %= x
		count += 1
	}

	return 0
}

func TestP26(t *testing.T) {

	periods := make([]int, 11)
	for i := 2; i <= 10; i++ {
		periods[i] = p26LenPeriod(i)
	}

	exp := []int{0, 0, 0, 1, 0, 0, 1, 6, 0, 1, 0}
	if !reflect.DeepEqual(periods, exp) {
		t.Error(periods)
	}

	max := 0
	maxv := 0
	for i := 2; i < 1000; i++ {
		p := p26LenPeriod(i)
		if max < p {
			max = p
			maxv = i
		}
	}

	if maxv != 983 {
		t.Error(maxv, max)
	}
}
