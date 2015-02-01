package natural

import "math/big"

func Factorial(n int) *big.Int {
	f := big.NewInt(1)
	for i := 1; i <= n; i++ {
		f.Mul(f, big.NewInt(int64(i)))
	}
	return f
}
