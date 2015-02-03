package euler

import (
	"go-euler/natural"
	"testing"
)

func TestP4(t *testing.T) {

	var z int

	max := 0
	for z1 := 999; z1 >= 100; z1-- {
		limit := z1
		if z1 < max/z1 {
			limit = max / z1
		}
		for z2 := 999; z2 >= limit; z2-- {
			z = z1 * z2
			if z > max && natural.DigitsPalindromic(natural.Digits(int64(z))) {
				max = z
			}
		}
	}

	if max != 906609 {
		t.Error(max)
	}
}
