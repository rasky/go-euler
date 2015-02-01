package euler

import (
	"euler/natural"
	"testing"
)

func p27PrimeSeqLen(a, b int, primes map[int32]bool, max int) int {

	isPrime := func(p int) bool {
		if p < 0 {
			return false
		}
		if p > max {
			panic("value too big, increase array")
		}
		if _, found := primes[int32(p)]; found {
			return true
		}
		return false
	}

	n := 0
	for isPrime(n*n + a*n + b) {
		n++
	}
	return n
}

func TestP27(t *testing.T) {

	const max = 100000

	primes := make(map[int32]bool, 1000)
	for p := range natural.Primes(max) {
		primes[p] = true
	}

	if p27PrimeSeqLen(1, 41, primes, max) != 40 {
		t.Error(1, 41, p27PrimeSeqLen(1, 41, primes, max))
	}
	if p27PrimeSeqLen(-79, 1601, primes, max) != 80 {
		t.Error(-79, 1601, p27PrimeSeqLen(-79, 1601, primes, max))
	}

	best, besta, bestb := 0, 0, 0
	for a := -1000; a <= 1000; a++ {
		for b := -1000; b <= 1000; b++ {
			l := p27PrimeSeqLen(a, b, primes, max)
			if best < l {
				best, besta, bestb = l, a, b
			}
		}
	}

	if besta*bestb != -59231 {
		t.Error(besta, bestb, best)
	}
}
