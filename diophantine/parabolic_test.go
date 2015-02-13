package diophantine

import (
	"math/big"
	"testing"
)

func TestParabolic(t *testing.T) {
	d := Equation{A: 8, B: -24, C: 18, D: 5, E: 7, F: 16}

	if !d.IsParabolic() {
		t.Error("not parabolic")
	}

	sol := d.Solve()
	defer sol.Finish()
	cnt := 0
	for c := range sol.Solutions {
		if !d.IsRoot(c.X, c.Y) {
			t.Error(c)
		}
		cnt++
		if cnt == 1000 {
			break
		}
	}
}

func newCoeffEq2(x0, x1, x2, y0, y1, y2 int64) CoeffEq2 {
	return CoeffEq2{
		X: [3]*big.Int{big.NewInt(x0), big.NewInt(x1), big.NewInt(x2)},
		Y: [3]*big.Int{big.NewInt(y0), big.NewInt(y1), big.NewInt(y2)},
	}
}

func TestParabolic2(t *testing.T) {
	eq := Equation{A: 9, B: 18, C: 9, D: 6, E: 8, F: 24}
	sol := eq.Solve()
	if sol.NumSolutions != SOLUTION_INFINITE || sol.codepath != "parabolic" {
		t.Error(sol)
	}
	if len(sol.Coeff) != 1 || !sol.Coeff[0].Equal(newCoeffEq2(12, 8, 18, -12, -6, -18)) {
		t.Error(sol)
	}
}

func TestParabolicParallel(t *testing.T) {
	d := Equation{A: 9, B: 18, C: 9, D: 5, E: 5, F: 6}
	sol := d.Solve()
	if sol.NumSolutions != SOLUTION_NONE || sol.codepath != "parabolic/parallel/noreal" {
		t.Error(sol)
	}

	d = Equation{A: 9, B: 18, C: 9, D: 15, E: 15, F: 6}
	sol = d.Solve()
	if sol.NumSolutions != SOLUTION_INFINITE || sol.codepath != "parabolic/parallel" {
		t.Error(sol)
	}
	if len(sol.Coeff) != 1 || !sol.Coeff[0].Equal(newCoeffEq2(0, 1, 0, -1, -1, 0)) {
		t.Error(sol)
	}

	d = Equation{A: 9, B: 18, C: 9, D: 20, E: 20, F: 4}
	sol = d.Solve()
	if sol.NumSolutions != SOLUTION_INFINITE || sol.codepath != "parabolic/parallel" {
		t.Error(sol)
	}
	if len(sol.Coeff) != 1 || !sol.Coeff[0].Equal(newCoeffEq2(0, 1, 0, -2, -1, 0)) {
		t.Error(sol)
	}

}
