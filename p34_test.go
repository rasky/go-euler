package euler

import (
	"testing"

	"code.google.com/p/intmath/intgr"
)

func TestP34(t *testing.T) {

	var fact [10]int

	fact[0] = 1
	for i := 1; i < 10; i++ {
		fact[i] = fact[i-1] * i
	}

	// Find upper bound
	i := 1
	for fact[9]*i > intgr.Pow(10, i)-1 {
		i += 1
	}
	upper := fact[9] * i

	total := 0
	for i := 10; i <= upper; i++ {
		j := i
		sum := 0
		for j != 0 {
			sum += fact[j%10]
			j /= 10
		}
		if sum == i {
			total += sum
		}
	}

	if total != 40730 {
		t.Error(total)
	}
}
