package euler

import (
	"go-euler/natural"
	"testing"
)

func TestP80(t *testing.T) {

	sqrt_sum := func(n int64, prec int) int {
		dig := natural.Sqrt(n, prec)
		sum := 0
		for d := range dig {
			if d == 0xFF {
				continue
			}
			sum += int(d)
		}
		return sum
	}

	if sqrt_sum(2, 100) != 475 {
		t.Error(sqrt_sum(2, 100))
	}

	sum := 0
	for i := int64(1); i <= 100; i++ {
		if !natural.IsSquare(i) {
			sum += sqrt_sum(i, 100)
		}
	}

	if sum != 40886 {
		t.Error(sum)
	}
}
