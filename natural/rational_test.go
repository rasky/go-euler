package natural

import "testing"

func TestRationalEqual(t *testing.T) {

	r1 := Rational{80, 81}
	r2 := Rational{0, 1}

	if r1.Equal(r2) {
		t.Error(r1, r2)
	}
}
