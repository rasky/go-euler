package euler

import (
	"euler/natural"
	"testing"
)

func TestP7(t *testing.T) {
	idx := 1
	var p int32
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
