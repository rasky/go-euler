package euler

import (
	"euler/natural"
	"testing"
)

func p27PrimeSeqLen(a, b int64) int64 {

	isPrime := func(p int64) bool {
		return p >= 0 && natural.IsPrime(p)
	}

	n := int64(1)
	for isPrime(n*n + a*n + b) {
		n++
	}
	return n - 1
}

func TestP27(t *testing.T) {

	if p27PrimeSeqLen(1, 41) != 40 {
		t.Error(1, 41, p27PrimeSeqLen(1, 41))
	}
	if p27PrimeSeqLen(-79, 1601) != 80 {
		t.Error(-79, 1601, p27PrimeSeqLen(-79, 1601))
	}

	best, besta, bestb := int64(0), int64(0), int64(0)
	for a := int64(-1000); a <= 1000; a++ {
		for b := int64(-1000); b <= 1000; b++ {
			l := p27PrimeSeqLen(a, b)
			if best < l {
				best, besta, bestb = l, a, b
			}
		}
	}

	if besta*bestb != -59231 {
		t.Error(besta, bestb, best)
	}
}
