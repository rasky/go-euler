package euler

import (
	"math"
	"math/big"
	"testing"

	"code.google.com/p/intmath/i64"
)

var maxMod = i64.Sqrt(math.MaxInt64)

func mulMod(a, b, mod int64) int64 {

	ax := big.NewInt(a)
	bx := big.NewInt(b)
	mx := big.NewInt(mod)
	ax.Mul(ax, bx)
	return ax.Mod(ax, mx).Int64()
}

func expModBig(base, exp, mod int64) int64 {
	res := int64(1)
	base = base % mod
	for exp > 0 {
		if exp&1 == 1 {
			res = mulMod(res, base, mod)
		}
		exp >>= 1
		base = mulMod(base, base, mod)
	}
	return res
}

func ExpMod(base, exp, mod int64) int64 {
	if mod > maxMod {
		return expModBig(base, exp, mod)
	}
	res := int64(1)
	base = base % mod
	for exp > 0 {
		if exp&1 == 1 {
			res = (res * base) % mod
		}
		exp >>= 1
		base = (base * base) % mod
	}
	return res
}

func TestP48(t *testing.T) {

	mod := int64(10000000000)

	sum := int64(0)
	for i := int64(1); i <= 1000; i++ {
		sum += ExpMod(i, i, mod)
		sum %= mod
		t.Log(sum)
	}

	if sum != 9110846700 {
		t.Error(sum)
	}
}
