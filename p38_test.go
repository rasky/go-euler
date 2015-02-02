package euler

import (
	"euler/natural"
	"testing"
)

func TestP38(t *testing.T) {
	max := int64(0)
	for i := int64(1); i < 99999; i++ {
		j := int64(1)
		prod := int64(0)
		proddig := make([]byte, 0, 9)
		for len(proddig) < 9 {
			prod = natural.Concat([]int64{prod, i * j})
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
