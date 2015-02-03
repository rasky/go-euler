package euler

import (
	"math/big"
	"testing"
)

func TestP57(t *testing.T) {

	count := 0
	num, den := big.NewInt(3), big.NewInt(2)

	for i := 1; i < 1000; i++ {
		den2 := new(big.Int).Add(num, den)
		num2 := new(big.Int).Add(den2, den)
		num, den = num2, den2
		if len(num.String()) > len(den.String()) {
			count += 1
		}

	}

	if count != 153 {
		t.Error(count)
	}
}
