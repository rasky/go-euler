package euler

import (
	"euler/natural"
	"testing"
)

func TestP3(t *testing.T) {
	var f int64
	for f = range natural.Factors(600851475143) {
	}

	if f != 6857 {
		t.Fail()
	}
}
