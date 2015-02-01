package euler

import (
	"euler/natural"
	"testing"
)

func TestP21(t *testing.T) {
	var d [10001]int

	for i := 1; i <= 10000; i++ {
		div := natural.Divisors(int64(i))
		sum := 0
		for _, j := range div {
			sum += j
		}
		d[i] = sum
	}

	if d[220] != 284 {
		t.Error(d[220], natural.Divisors(220))
	}
	if d[284] != 220 {
		t.Error(d[284], natural.Divisors(284))
	}

	sum := 0
	for i := 1; i <= 10000; i++ {
		a := d[i]
		if a <= 10000 && a != i && d[a] == i {
			sum += i
		}
	}

	if sum != 31626 {
		t.Error(sum)
	}
}
