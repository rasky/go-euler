package diophantine

import "testing"

func TestHyperbolic(t *testing.T) {

	d := Equation{A: 0, B: 2, C: 0, D: 5, E: 56, F: 7}

	if !d.IsHyperbolic() {
		t.Error("not hyperbolic")
	}

	sol := d.Solve()
	defer sol.Finish()
	cnt := 0
	for c := range sol.Solutions {
		if !d.IsRoot(c.X, c.Y) {
			t.Error(c)
		}
		cnt++
	}

	if cnt != 8 {
		t.Error(cnt)
	}
}
