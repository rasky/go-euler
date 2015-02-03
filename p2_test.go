package euler

import (
	"go-euler/natural"
	"testing"
)

func TestP2(t *testing.T) {
	sum := 0
	for i := range natural.Fibo() {
		if i >= 4000000 {
			break
		}
		if i%2 == 0 {
			sum += i
		}
	}

	if sum != 4613732 {
		t.Fail()
	}
}
