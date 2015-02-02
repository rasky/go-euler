package euler

import (
	"euler/natural"

	"code.google.com/p/intmath/intgr"

	"testing"
)

func p32NumDigits(a int) int {
	switch {
	case a < 10:
		return 1
	case a < 100:
		return 2
	case a < 1000:
		return 3
	case a < 10000:
		return 4
	case a < 100000:
		return 5
	case a < 1000000:
		return 6
	case a < 10000000:
		return 7
	case a < 100000000:
		return 8
	case a < 1000000000:
		return 9
	case a < 10000000000:
		return 10
	}
	panic("unreachable")
}

// Return c=a*b if "a*b=c" is pandigital; return 0 otherwise
func p32PanDigitalProduct(a, b int) int {
	c := a * b

	if natural.PandigitalMulti([]int{a, b, c}, 9, false) {
		return c
	}
	return 0
}

func TestP32(t *testing.T) {

	if p32PanDigitalProduct(39, 186) != 7254 {
		t.Fail()
	}

	pandig := make(map[int]interface{})

	sum := 0
	for a := 2; a < 999; a++ {
		ad := p32NumDigits(a)
		bdmin := (9 - ad) / 2
		bmin := intgr.Pow(10, bdmin-1)
		for b := bmin; b < bmin*10; b++ {
			c := p32PanDigitalProduct(a, b)
			if c != 0 {
				if _, found := pandig[c]; !found {
					sum += c
					pandig[c] = nil
				}
			}
		}
	}

	if sum != 45228 {
		t.Error(sum)
	}
}
