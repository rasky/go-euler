package euler

import (
	"euler/natural"
	"math/big"
	"testing"
)

func TestP25(t *testing.T) {

	limit := big.NewInt(10)
	limit = limit.Exp(limit, big.NewInt(1000-1), nil)

	idx := 1 // consider two-times "1" at beginning
	for f := range natural.FiboBig() {
		idx += 1
		if f.Cmp(limit) >= 0 {
			break
		}
	}

	if idx != 4782 {
		t.Error(idx)
	}
}
