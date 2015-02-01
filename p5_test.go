package euler

import (
	"euler/natural"
	"testing"

	"code.google.com/p/intmath/i32"
)

func TestP5(t *testing.T) {

	allfactors := make(map[int32]int32)

	for i := 2; i <= 20; i++ {
		factors := make(map[int32]int32)
		for f := range natural.Factors(int64(i)) {
			factors[f] += 1
		}
		for k, v := range factors {
			if allfactors[k] < v {
				allfactors[k] = v
			}
		}
	}

	total := int32(1)
	for k, v := range allfactors {
		total *= i32.Pow(k, v)
	}

	if total != 232792560 {
		t.Error(total)
	}
}
