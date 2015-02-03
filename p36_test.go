package euler

import (
	"go-euler/natural"
	"testing"
)

func TestP36(t *testing.T) {

	sum := int64(0)

	for i := int64(1); i < 1000000; i++ {
		if !natural.DigitsPalindromic(natural.DigitsWithBase(i, 10)) {
			continue
		}
		if !natural.DigitsPalindromic(natural.DigitsWithBase(i, 2)) {
			continue
		}
		sum += i
	}

	if sum != 872187 {
		t.Error(sum)
	}
}
