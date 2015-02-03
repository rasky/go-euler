package euler

import (
	"go-euler/natural"
	"testing"
)

func TestP62(t *testing.T) {

	dig_memo := make(map[int64][10]byte)
	digits := func(x int64) ([10]byte, int64) {
		if dig, found := dig_memo[x]; found {
			return dig, -1
		}
		var dig [10]byte
		sum := int64(0)
		for _, d := range natural.Digits(x) {
			sum += int64(d)
			dig[d]++
		}
		dig_memo[x] = dig
		return dig, sum
	}

	same := func(x, y [10]byte) bool {
		for i := 0; i < 10; i++ {
			if x[i] != y[i] {
				return false
			}
		}
		return true
	}

	nextpow := int64(1000)
	numperm := make(map[int64]int, 1024)
	bysum := make(map[int64][]int64, 1024)

	var solution int64

Loop:
	for i := int64(500); i < 100000; i++ {
		if i >= nextpow {
			numperm = make(map[int64]int, 1024)
			bysum = make(map[int64][]int64, 1024)
			nextpow *= 10
		}
		n := i * i * i
		dig, sum := digits(n)
		found := false
		for _, m := range bysum[sum] {
			dig2, _ := digits(m)
			if same(dig, dig2) {
				numperm[m]++
				if numperm[m] == 5 {
					solution = m
					break Loop
				}
				found = true
				break
			}
		}
		if !found {
			numperm[n] = 1
		}
		bysum[sum] = append(bysum[sum], n)
	}

	if solution != 127035954683 {
		t.Error(solution)
	}

}
