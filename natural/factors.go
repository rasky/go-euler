package natural

import "code.google.com/p/intmath/i64"

type factors map[int32]int

func Factors(x int64) chan int32 {

	ch := make(chan int32)

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
		for y := range Primes(int32(limit)) {
			for (x % int64(y)) == 0 {
				x = x / int64(y)
				ch <- y
			}
		}
		if x != 1 {
			ch <- int32(x)
		}
	}()

	return ch
}

func Factorize(x int64) map[int32]int32 {

	factors := make(map[int32]int32, 32)
	for f := range Factors(x) {
		factors[f] += 1
	}

	return factors
}
