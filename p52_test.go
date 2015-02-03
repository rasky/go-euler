package euler

import (
	"go-euler/natural"
	"testing"
)

func TestP52(t *testing.T) {
	var i int64
	for i = 1; i < 1000000; i++ {
		dig := natural.Digits(i)
		// early-exit
		if !natural.IsPermutation(dig, natural.Digits(i*2)) {
			continue
		}
		perms := [][]byte{
			natural.Digits(i * 3),
			natural.Digits(i * 4),
			natural.Digits(i * 5),
			natural.Digits(i * 6),
		}
		if natural.IsPermutationMulti(dig, perms) {
			break
		}
	}

	if i != 142857 {
		t.Error(i)
	}
}
