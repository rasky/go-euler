package euler

import (
	"euler/natural"
	"math/big"
	"testing"
)

func TestP56(t *testing.T) {

	max := 0
	for a := int64(2); a < 100; a++ {
		n := big.NewInt(int64(a))
		n1 := big.NewInt(int64(a))
		b := 1
		for b < 100 {
			sum := natural.BigSumDigits(n)
			if max < sum {
				max = sum
				t.Log(max, n)
			}
			n.Mul(n, n1)
			b++
		}
	}

	if max != 972 {
		t.Error(max)
	}
}
