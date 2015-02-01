package euler

import "testing"

func TestP9(t *testing.T) {

	var found bool
	var a, b, c int
Loop:
	for c = 1000; c > 0; c-- {
		for b = 1000 - c; b > 0; b-- {
			a = 1000 - c - b
			if a > 0 && a*a+b*b == c*c {
				found = true
				break Loop
			}
		}
	}

	if !found {
		t.Fail()
	}

	if a*b*c != 31875000 {
		t.Error(a * b * c)
	}
}
