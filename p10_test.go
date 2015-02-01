package euler

import (
	"euler/natural"
	"testing"
)

func TestP10(t *testing.T) {

	total := int64(0)
	for p := range natural.Primes(2000000) {
		total += int64(p)
	}

	if total != 142913828922 {
		t.Error(total)
	}

}
