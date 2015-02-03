package euler

import (
	"go-euler/natural"
	"testing"
)

func TestP49(t *testing.T) {

	found := make([]int64, 0)

	for px := range natural.Primes(10000) {
		p := int64(px)
		if p < 1000 {
			continue
		}
		p1 := p + 3330
		p2 := p1 + 3330

		if p2 >= 10000 {
			break
		}

		dig := natural.Digits(p)
		if natural.IsPermutation(dig, natural.Digits(p1)) &&
			natural.IsPermutation(dig, natural.Digits(p2)) &&
			natural.IsPrime(p1) &&
			natural.IsPrime(p2) {
			found = append(found, natural.Concat([]int64{p, p1, p2}))
		}
	}

	if found[0] != 148748178147 || found[1] != 296962999629 {
		t.Error(found)
	}
}
