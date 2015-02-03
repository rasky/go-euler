package euler

import (
	"math/big"

	"testing"
)

func TestP63(t *testing.T) {

	count := 0
	base := big.NewInt(1)
	limit := big.NewInt(10)

	for n := int64(1); ; n++ {

		prevcount := count

		var j int64
		for j = 1; ; j++ {
			p := new(big.Int).Exp(big.NewInt(j), big.NewInt(n), nil)
			if p.Cmp(base) < 0 {
				continue
			}
			if p.Cmp(limit) < 0 {
				t.Log(j, n, p)
				count++
			} else {
				break
			}
		}

		if prevcount == count {
			break
		}

		base = new(big.Int).Mul(base, big.NewInt(10))
		limit = new(big.Int).Mul(limit, big.NewInt(10))
	}

	if count != 49 {
		t.Error(count)
	}
}
