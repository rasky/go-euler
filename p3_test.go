package euler

import (
	"go-euler/natural"
	"testing"
)

func TestP3(t *testing.T) {
	factors := natural.Factorize(600851475143)

	max := int64(0)
	for f, _ := range factors {
		if max < f {
			max = f
		}
	}

	if max != 6857 {
		t.Fail()
	}
}
