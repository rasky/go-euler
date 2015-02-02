package euler

import (
	"euler/natural"
	"testing"
)

func TestP38(t *testing.T) {
	max := 0
	for i := 1; i < 99999; i++ {
		j := 1
		prod := 0
		proddig := make([]byte, 0, 9)
		for len(proddig) < 9 {
			prod = natural.Concat([]int{prod, i * j})
			proddig = append(proddig, natural.Digits(i*j)...)
			j += 1
		}
		if natural.PandigitalDigits(proddig, 9, false) {
			if max < prod {
				max = prod
			}
		}
	}

	if max != 932718654 {
		t.Error(max)
	}
}
