package natural

import "math/big"

func Factorial(n int) *big.Int {
	f := big.NewInt(1)
	for i := 1; i <= n; i++ {
		f.Mul(f, big.NewInt(int64(i)))
	}
	return f
}

func FactorialArray(n int64) []*big.Int {
	fact := make([]*big.Int, n)
	fact[0] = big.NewInt(1)
	for i := int64(1); i < n; i++ {
		fact[i] = big.NewInt(0).Mul(fact[i-1], big.NewInt(i))
	}
	return fact
}
