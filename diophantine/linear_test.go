package diophantine

import "testing"

func TestLinear(t *testing.T) {

	d := Equation{D: 10, E: 84, F: 16}
	if !d.IsLinear() {
		t.Error("not linear", d)
	}

	sol := d.Solve()
	defer sol.Finish()

	if sol.NumSolutions != SOLUTION_INFINITE {
		t.Error(sol.NumSolutions)
	}

	for i := 0; i < 50; i++ {
		p := <-sol.Solutions
		if !d.IsRoot(p.X, p.Y) {
			t.Error(p)
		}
	}
}
