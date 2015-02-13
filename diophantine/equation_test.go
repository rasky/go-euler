package diophantine

import "testing"

func TestBasicGcd(t *testing.T) {

	eq := Equation{A: 4, B: 12, C: 8, D: -6, E: -2, F: 9}
	sol := eq.Solve()
	if sol.NumSolutions != 0 || sol.codepath != "nogcd" {
		t.Error("no solutions should be found")
	}
}

func TestBasicNoMod(t *testing.T) {

	eq := Equation{A: 4, B: 12, C: 9, D: -6, E: -9, F: 9}
	sol := eq.Solve()
	if sol.NumSolutions != 0 || sol.codepath != "nomod" {
		t.Error("no mod should be found")
	}
}
