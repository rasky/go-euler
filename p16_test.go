package euler

import (
	"math/big"
	"testing"
)

func p16SumDigits(b *big.Int) int {
	num := 0
	for _, s := range b.String() {
		num += int(s - '0')
	}
	return num
}

func TestP16(t *testing.T) {

	var b big.Int
	b.Exp(big.NewInt(2), big.NewInt(15), nil)
	if p16SumDigits(&b) != 26 {
		t.Error(p16SumDigits(&b))
	}

	b.Exp(big.NewInt(2), big.NewInt(1000), nil)
	if p16SumDigits(&b) != 1366 {
		t.Error(p16SumDigits(&b))
	}
}
