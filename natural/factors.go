package natural

import "code.google.com/p/intmath/i64"

func Factorize(x int64) map[int64]int {
	factors := make(map[int64]int, 32)

	for x%2 == 0 {
		x /= 2
		factors[2]++
	}
	for x%3 == 0 {
		x /= 3
		factors[3]++
	}

	for y := range Primes(i64.Sqrt(x)) {
		for (x % y) == 0 {
			x = x / y
			factors[y]++
		}
	}
	if x != 1 {
		factors[x]++
	}
	return factors
}
