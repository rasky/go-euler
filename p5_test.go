package euler

import (
	"euler/natural"
	"testing"

	"code.google.com/p/intmath/i64"
)

func TestP5(t *testing.T) {

	allfactors := make(map[int64]int)

	for i := int64(2); i <= 20; i++ {
		factors := natural.Factorize(i)
		for k, v := range factors {
			if allfactors[k] < v {
				allfactors[k] = v
			}
		}
	}

	total := int64(1)
	for k, v := range allfactors {
		total *= i64.Pow(k, int64(v))
	}

	if total != 232792560 {
		t.Error(total)
	}
}
