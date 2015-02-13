package diophantine

import "math/big"

func (eq Equation) IsLinear() bool {
	return eq.A == 0 && eq.B == 0 && eq.C == 0
}

func (eq Equation) solveLinear() (sol Solution) {

	if eq.D == 0 && eq.E == 0 {
		if eq.F != 0 {
			sol.NumSolutions = SOLUTION_NONE
		} else {
			sol.NumSolutions = SOLUTION_ALL
		}
		return
	} else if eq.D == 0 {
		if eq.F%eq.E != 0 {
			return
		}
		by := big.NewInt(-eq.F / eq.E)
		sol.Coeff = append(sol.Coeff, CoeffEq2{
			X: [3]*big.Int{bzero, bone, bzero},
			Y: [3]*big.Int{by, bzero, bzero},
		})
		sol.generate(GEN_ALL_INTEGERS, func(n int64) (*big.Int, *big.Int) {
			return big.NewInt(n), by
		})
		return
	} else if eq.E == 0 {
		if eq.F%eq.D != 0 {
			return
		}
		bx := big.NewInt(-eq.F / eq.D)
		sol.Coeff = append(sol.Coeff, CoeffEq2{
			X: [3]*big.Int{bx, bzero, bzero},
			Y: [3]*big.Int{bzero, bone, bzero},
		})
		sol.generate(GEN_ALL_INTEGERS, func(n int64) (*big.Int, *big.Int) {
			return bx, big.NewInt(n)
		})
		return
	}

	g := gcd(eq.D, eq.E)
	if eq.F%g != 0 {
		return
	}

	d, e, f := eq.D/g, eq.E/g, eq.F/g
	exgcd := extendedGCD(eq.D, eq.E)
	fu := -exgcd.U * f
	fv := -exgcd.V * f
	sol.Coeff = append(sol.Coeff, CoeffEq2{
		X: [3]*big.Int{big.NewInt(fu), big.NewInt(e), bzero},
		Y: [3]*big.Int{big.NewInt(fv), big.NewInt(-d), bzero},
	})
	sol.codepath = "linear"

	sol.generate(GEN_ALL_INTEGERS, func(n int64) (*big.Int, *big.Int) {
		return big.NewInt(e*n + fu), big.NewInt(-d*n + fv)
	})

	return
}
