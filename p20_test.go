package euler

import (
	"euler/natural"
	"testing"
)

func TestP20(t *testing.T) {

	f := natural.Factorial(10)
	if natural.BigSumDigits(f) != 27 {
		t.Error(natural.BigSumDigits(f))
	}

	f = natural.Factorial(100)
	if natural.BigSumDigits(f) != 648 {
		t.Error(natural.BigSumDigits(f))
	}
}
