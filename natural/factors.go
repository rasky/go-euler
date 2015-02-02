package natural

import "code.google.com/p/intmath/i64"

func Factors(x int64) chan int64 {

	ch := make(chan int64)

	go func() {
		defer close(ch)

		for x%2 == 0 {
			x /= 2
			ch <- 2
		}
		for x%3 == 0 {
			x /= 3
			ch <- 3
		}
		limit := i64.Sqrt(x)
		for y := range Primes(limit) {
			for (x % y) == 0 {
				x = x / y
				ch <- y
			}
		}
		if x != 1 {
			ch <- x
		}
	}()

	return ch
}

func Factorize(x int64) map[int64]int {

	factors := make(map[int64]int, 32)
	for f := range Factors(x) {
		factors[f] += 1
	}

	return factors
}
