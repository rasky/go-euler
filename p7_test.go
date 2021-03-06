package euler

import (
	"go-euler/natural"
	"testing"
)

func TestP7(t *testing.T) {
	idx := 1
	var p int64
	for p = range natural.Primes(1000000) {
		if idx == 10001 {
			break
		}
		idx += 1
	}

	if p != 104743 {
		t.Error(p)
	}
}
