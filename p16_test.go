package euler

import (
	"euler/natural"
	"math/big"
	"testing"
)

func TestP16(t *testing.T) {

	var b big.Int
	b.Exp(big.NewInt(2), big.NewInt(15), nil)
	if natural.BigSumDigits(&b) != 26 {
		t.Error(natural.BigSumDigits(&b))
	}

	b.Exp(big.NewInt(2), big.NewInt(1000), nil)
	if natural.BigSumDigits(&b) != 1366 {
		t.Error(natural.BigSumDigits(&b))
	}
}
