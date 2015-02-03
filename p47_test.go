package euler

import (
	"go-euler/natural"
	"testing"
)

func TestP47(t *testing.T) {

	consecFactors := func(search int) int64 {
		count := 0
		n := int64(0)
		for n = 2; n < 1000000; n++ {
			factors := natural.Factorize(n)
			if len(factors) == search {
				count += 1
				if count == search {
					break
				}
			} else {
				count = 0
			}
		}
		return n - int64(search) + 1
	}

	c2 := consecFactors(2)
	c3 := consecFactors(3)
	c4 := consecFactors(4)

	if c2 != 14 || c3 != 644 || c4 != 134043 {
		t.Error(c2, c3, c4)
	}
}
