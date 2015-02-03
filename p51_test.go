package euler

import (
	"go-euler/natural"
	"testing"
)

func p51CheckReplacement(dig []byte, d byte, family_size int) []int64 {

	pos := make([]int, 0, 8)
	for idx, d2 := range dig {
		if d == d2 {
			pos = append(pos, idx)
		}
	}

	last := len(dig) - 1
	matches := make([]int64, 1, 8)
	matches[0] = natural.FromDigits(dig)
	match := 0
	for i := byte(0); i <= 9; i++ {
		if i == d {
			continue
		}
		for _, j := range pos {
			dig[j] = i
		}
		if dig[last] == 0 {
			continue
		}
		p := natural.FromDigits(dig)
		if natural.IsPrime(p) {
			match += 1
			matches = append(matches, p)
			if match >= family_size {
				return matches
			}
		}
	}
	return []int64{}
}

func TestP51(t *testing.T) {

	var matches []int64

	const family_size = 7

Loop:
	for p := range natural.Primes(1000000) {
		if p < 10000 {
			continue
		}

		dig := natural.Digits(p)
		dig2 := make([]byte, len(dig))
		seen := make([]bool, 10)
		for _, d := range dig {
			if !seen[d] {
				seen[d] = true
				copy(dig2, dig)
				if matches = p51CheckReplacement(dig2, d, family_size); len(matches) > 0 {
					break Loop
				}
			}
		}
	}

	if len(matches) == 0 {
		t.Fail()
	}

	min := matches[0]
	for _, m := range matches {
		if min > m {
			min = m
		}
	}

	if min != 121313 {
		t.Error(min, matches)
	}
}
