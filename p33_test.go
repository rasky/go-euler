package euler

import (
	"euler/natural"
	"testing"
)

func TestP33(t *testing.T) {

	product := natural.Rational{1, 1}
	count := 0

	for a := 10; a < 100; a++ {
		for b := a + 1; b < 100; b++ {

			f := natural.Rational{a, b}

			a1 := a / 10
			a2 := a % 10
			b1 := b / 10
			b2 := b % 10

			if a1 == b1 && b2 != 0 && f.Equal(natural.Rational{a2, b2}) {
				product = product.Mul(f)
				count += 1
			} else if a1 == b2 && f.Equal(natural.Rational{a2, b1}) {
				product = product.Mul(f)
				count += 1
			} else if a2 == b1 && b2 != 0 && f.Equal(natural.Rational{a1, b2}) {
				product = product.Mul(f)
				count += 1
			} else if a2 == b2 && f.Equal(natural.Rational{a1, b1}) {
				if a2 != 0 {
					product = product.Mul(f)
					count += 1
				}
			}
		}
	}

	if count != 4 {
		t.Error(count)
	}

	product = product.Reduce()
	if product.Den != 100 {
		t.Error(product)
	}
}
