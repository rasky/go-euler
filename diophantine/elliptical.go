package diophantine

import "math/big"

func (eq Equation) IsElliptical() bool {
	return eq.B*eq.B-4*eq.A*eq.C < 0
}

func (eq Equation) solveElliptical() (sol Solution) {

	a := eq.B*eq.B - 4*eq.A*eq.C
	b := 2 * (eq.B*eq.E - 2*eq.C*eq.D)
	c := eq.E*eq.E - 4*eq.C*eq.F

	if b*b-4*a*c < 0 {
		sol.codepath = "elliptical/noroots"
		return
	}

	sq := isqrt(b*b - 4*a*c)
	min := (-b - sq) / (2 * a)
	max := (-b + sq) / (2 * a)
	if min > max {
		min, max = max, min
	}

	x := min - 1
	sign := int64(-1)
	next := func() bool {
		if sign > 0 {
			sign = -sign
			return true
		} else {
			sign = -sign
			x++
			return (x <= max)
		}
	}

	sol.codepath = "elliptical"
	sol.generate(GEN_ALL_INTEGERS, func(n int64) (*big.Int, *big.Int) {
		for next() {
			t1 := -(eq.B*x + eq.E)
			t2 := t1*t1 - 4*eq.C*(eq.A*x*x+eq.D*x+eq.F)
			t2sq := isqrt(t2)
			if t2 != t2sq*t2sq {
				continue
			}
			num := t1 + t2sq*sign
			den := 2 * eq.C
			if num%den != 0 {
				continue
			}

			y := num / den
			return big.NewInt(x), big.NewInt(y)
		}
		return nil, nil
	})
	return
}
