package euler

import "testing"

func TestP1(t *testing.T) {

	sum := 0
	for i := 0; i < 1000; i++ {
		if i%3 == 0 || i%5 == 0 {
			sum += i
		}
	}

	if sum != 233168 {
		t.Fail()
	}
}
