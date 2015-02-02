package euler

import (
	"euler/natural"
	"testing"
)

func TestP40(t *testing.T) {

	dig := make([]byte, 1, 1100000)
	for i := 1; len(dig) <= 1000000; i++ {
		idig := natural.Digits(i)
		natural.Reverse(idig)
		dig = append(dig, idig...)
	}

	if dig[12] != 1 {
		t.Error(dig[12])
	}

	ret := dig[1] * dig[10] * dig[100] * dig[1000] * dig[10000] * dig[100000] * dig[1000000]
	if ret != 210 {
		t.Error(ret)
	}
}
