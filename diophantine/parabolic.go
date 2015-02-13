package diophantine

import "math/big"

func (eq Equation) IsParabolic() bool {
	return eq.B*eq.B-4*eq.A*eq.C == 0
}

func (eq Equation) solveParabolic() (sol Solution) {

	g := gcd(eq.A, eq.C)
	if eq.A < 0 {
		g = -g
	}
	a := eq.A / g
	c := eq.C / g
	sqa := isqrt(a)
	sqc := isqrt(c)
	if eq.A*eq.B < 0 {
		sqc = -sqc
	}

	CDmAE := sqc*eq.D - sqa*eq.E
	if CDmAE == 0 {
		// Two parallel lines
		au := sqa * g
		bu := eq.D
		cu := sqa * eq.F
		b24acu := bu*bu - 4*au*cu
		if b24acu < 0 {
			sol.codepath = "parabolic/parallel/noreal"
			return
		}

		var solutions []Solution

		for _, u := range []int64{-bu + isqrt(b24acu), -bu - isqrt(b24acu)} {
			if u%(2*au) != 0 {
				continue
			}
			u /= 2 * au

			d := Equation{D: sqa, E: sqc, F: -u}
			dsol := d.Solve()
			solutions = append(solutions, dsol)
			sol.Coeff = append(sol.Coeff, dsol.Coeff...)
		}

		flip := false
		sol.codepath = "parabolic/parallel"
		sol.generate(GEN_ALL_INTEGERS, func(n int64) (*big.Int, *big.Int) {
			flip = !flip
			if flip {
				c := <-solutions[0].Solutions
				return c.X, c.Y
			} else {
				c := <-solutions[1].Solutions
				return c.X, c.Y
			}
		})
		return
	}

	// A parabola
	au := sqa * g
	bu := eq.D
	cu := sqa * eq.F
	max := CDmAE
	if max < 0 {
		max = -max
	}

	ax := big.NewInt(sqc * g * (sqa*eq.E - sqc*eq.D))
	ay := big.NewInt(sqa * g * (sqc*eq.D - sqa*eq.E))

	for u := int64(0); u < max; u++ {
		if (au*u*u+bu*u+cu)%max == 0 {
			sol.Coeff = append(sol.Coeff,
				CoeffEq2{
					X: [3]*big.Int{
						big.NewInt(-(sqc*g*u*u + eq.E*u + sqc*eq.F) / CDmAE),
						big.NewInt(+(eq.E + 2*sqc*g*u)),
						ax,
					},
					Y: [3]*big.Int{
						big.NewInt(+(sqa*g*u*u + eq.D*u + sqa*eq.F) / CDmAE),
						big.NewInt(-(eq.D + 2*sqa*g*u)),
						ay,
					},
				})
		}
	}
	if len(sol.Coeff) == 0 {
		sol.codepath = "parabolic/nocoeffs"
		return
	}

	idx := 0
	cnt := int64(0)
	t := int64(0)
	next := func() {
		idx++
		if idx == len(sol.Coeff) {
			idx = 0
			cnt++
			t = (cnt + 1) / 2
			if cnt%2 == 0 {
				t = -t
			}
		}
	}

	sol.codepath = "parabolic"
	sol.generate(GEN_ALL_INTEGERS, func(n int64) (*big.Int, *big.Int) {
		bt := big.NewInt(t)
		bt2 := new(big.Int).Mul(bt, bt)

		xta := new(big.Int).Mul(bt2, sol.Coeff[idx].X[2])
		yta := new(big.Int).Mul(bt2, sol.Coeff[idx].Y[2])
		xtb := new(big.Int).Mul(sol.Coeff[idx].X[1], bt)
		ytb := new(big.Int).Mul(sol.Coeff[idx].Y[1], bt)

		bx := new(big.Int).Add(xta, xtb)
		bx.Add(bx, sol.Coeff[idx].X[0])
		by := new(big.Int).Add(yta, ytb)
		by.Add(by, sol.Coeff[idx].Y[0])

		next()

		return bx, by
	})

	return
}
