package diophantine

import "testing"

func TestElliptical(t *testing.T) {

	d := Equation{A: 42, B: 8, C: 15, D: 23, E: 17, F: -4915}

	if !d.IsElliptical() {
		t.Error("not elliptical")
	}

	sol := d.Solve()
	cnt := 0
	for c := range sol.Solutions {
		if !d.IsRoot(c.X, c.Y) {
			t.Error(c)
		}
		t.Log(c)
		cnt++
	}

	if cnt != 1 {
		t.Error(cnt)
	}
}

func TestEllipticalNoRoots(t *testing.T) {

	eq := Equation{A: 4, B: 12, C: 10, D: -6, E: -9, F: 9}
	sol := eq.Solve()
	if sol.NumSolutions != 0 || sol.codepath != "elliptical/noroots" {
		t.Error("no roots should be found")
	}
}
