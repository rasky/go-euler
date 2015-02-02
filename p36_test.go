package euler

import (
	"euler/natural"
	"testing"
)

func p36Palindromic(dig []byte) bool {
	for i := 0; i < len(dig)/2; i++ {
		if dig[i] != dig[len(dig)-i-1] {
			return false
		}
	}
	return true
}

func TestP36(t *testing.T) {

	sum := 0

	for i := 1; i < 1000000; i++ {
		if !p36Palindromic(natural.DigitsWithBase(i, 10)) {
			continue
		}
		if !p36Palindromic(natural.DigitsWithBase(i, 2)) {
			continue
		}
		sum += i
	}

	if sum != 872187 {
		t.Error(sum)
	}
}
