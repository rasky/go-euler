package euler

import (
	"testing"

	"code.google.com/p/intmath/intgr"
)

func p30DigitPowers(pow int) []int {

	found := make([]int, 0)

	for i := 10; i < 1000000; i++ {
		x := i
		sum := 0
		for x != 0 {
			dig := x % 10
			sum += intgr.Pow(dig, pow)
			x /= 10
		}
		if sum == i {
			found = append(found, i)
		}
	}
	return found
}

func TestP30(t *testing.T) {

	sum := 0
	for _, v := range p30DigitPowers(4) {
		sum += v
	}
	if sum != 19316 {
		t.Error(sum)
	}

	sum = 0
	for _, v := range p30DigitPowers(5) {
		sum += v
	}
	if sum != 443839 {
		t.Error(sum)
	}

}
