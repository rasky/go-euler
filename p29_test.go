package euler

import (
	"math/big"
	"testing"
)

func TestP29(t *testing.T) {

	power := make(map[string]interface{})

	one := big.NewInt(1)
	top := big.NewInt(100)

	for a := big.NewInt(2); a.Cmp(top) <= 0; a.Add(a, one) {
		for b := big.NewInt(2); b.Cmp(top) <= 0; b.Add(b, one) {
			res := big.NewInt(0).Exp(a, b, nil).String()
			power[res] = nil
		}
	}

	if len(power) != 9183 {
		t.Error(power)
	}
}
