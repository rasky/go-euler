package euler

import (
	"euler/natural"
	"testing"
)

func TestP41(t *testing.T) {
	max := int64(0)

	digits := []byte{1, 2, 3, 4, 5, 6, 7}
	loop := true
	for loop {
		p := natural.FromDigits(digits)
		if p > max && natural.IsPrime(p) {
			max = p
		}
		loop = natural.NextPermutation(digits)
	}

	if max != 7652413 {
		t.Error(max)
	}
}
