package diophantine

import "math/big"

func (eq Equation) IsHyperbolic() bool {
	return eq.A == 0 && eq.C == 0 && eq.B != 0
}

func (eq Equation) solveHyperbolic() (sol Solution) {
	DEmBF := eq.D*eq.E - eq.B*eq.F
	if DEmBF == 0 {
		// Two lines parallel to x/y axes respectively
		if eq.E%eq.B == 0 {
			bx := big.NewInt(-eq.E * eq.B)
			sol.generate(GEN_ALL_INTEGERS, func(n int64) (*big.Int, *big.Int) {
				return bx, big.NewInt(n)
			})
		} else {
			by := big.NewInt(-eq.D * eq.B)
			sol.generate(GEN_ALL_INTEGERS, func(n int64) (*big.Int, *big.Int) {
				return big.NewInt(n), by
			})
		}
		return
	}

	d := int64(1)
	sol.generate(GEN_ALL_INTEGERS_NONZERO, func(n int64) (*big.Int, *big.Int) {
		if d > DEmBF {
			return nil, nil
		}
		for DEmBF%d != 0 || (d-eq.E)%eq.B != 0 {
			d++
		}
		bx := big.NewInt((d - eq.E) / eq.B)
		by := big.NewInt((DEmBF/d - eq.D) / eq.B)
		if d >= 0 {
			d = -d
		} else {
			d = -d
			d++
		}
		return bx, by
	})
	return
}
