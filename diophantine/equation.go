package diophantine

import (
	"fmt"
	"math/big"
	"strings"
)

// A structure describing a diophantine equation in two variables:
//    Ax^2 + Bxy + Cy^2 + Dx + Ey + F = 0
type Equation struct {
	A, B, C, D, E, F int64
}

type Coords struct {
	X, Y *big.Int
}

var bzero *big.Int = big.NewInt(0)
var bone *big.Int = big.NewInt(1)
var bmone *big.Int = big.NewInt(-1)

func formatTerm(n int64) string {
	switch n {
	case 1:
		return "+"
	case -1:
		return "-"
	default:
		return fmt.Sprintf("%+d", n)
	}
}

func formatTermBig(n *big.Int) string {
	if n.Cmp(bone) == 0 {
		return "+"
	}
	if n.Cmp(bmone) == 0 {
		return "-"
	}
	return fmt.Sprintf("%+d", n)
}

func (eq Equation) String() string {
	terms := make([]string, 0)
	if eq.A != 0 {
		terms = append(terms, fmt.Sprintf("%sx²", formatTerm(eq.A)))
	}
	if eq.B != 0 {
		terms = append(terms, fmt.Sprintf("%sxy", formatTerm(eq.B)))
	}
	if eq.C != 0 {
		terms = append(terms, fmt.Sprintf("%sy²", formatTerm(eq.C)))
	}
	if eq.D != 0 {
		terms = append(terms, fmt.Sprintf("%sx", formatTerm(eq.D)))
	}
	if eq.E != 0 {
		terms = append(terms, fmt.Sprintf("%sy", formatTerm(eq.E)))
	}
	if eq.F != 0 {
		terms = append(terms, fmt.Sprintf("%s", formatTerm(eq.F)))
	}

	eqs := strings.Join(terms, "")
	if eqs[0] == '+' {
		return "[" + eqs[1:] + "=0]"
	} else {
		return "[" + eqs + "=0]"
	}
}

func (eq Equation) IsRoot(X, Y *big.Int) bool {
	X2 := new(big.Int).Mul(X, X)
	XY := new(big.Int).Mul(X, Y)
	Y2 := new(big.Int).Mul(Y, Y)

	X2.Mul(X2, big.NewInt(eq.A))
	XY.Mul(XY, big.NewInt(eq.B))
	Y2.Mul(Y2, big.NewInt(eq.C))

	DX := new(big.Int).Mul(X, big.NewInt(eq.D))
	EY := new(big.Int).Mul(Y, big.NewInt(eq.E))

	V := big.NewInt(eq.F)
	V.Add(V, X2)
	V.Add(V, XY)
	V.Add(V, Y2)
	V.Add(V, DX)
	V.Add(V, EY)
	return V.Cmp(bzero) == 0
}

func (eq Equation) solveMod(mod int64) bool {
	for x := int64(0); x < mod; x++ {
		z := eq.A*x*x + eq.D*x + eq.F
		t := eq.B*x + eq.E
		for y := int64(0); y < mod; y++ {
			if (z+y*(t+eq.C*y))%mod == 0 {
				return false
			}
		}
	}
	return true
}

func (eq Equation) Solve() (sol Solution) {
	g := gcd(eq.A, eq.B, eq.C, eq.D, eq.E)
	if eq.F%g != 0 {
		sol.codepath = "nogcd"
		return
	}
	eq.A /= g
	eq.B /= g
	eq.C /= g
	eq.D /= g
	eq.E /= g
	eq.F /= g

	if eq.IsLinear() {
		return eq.solveLinear()
	}

	if eq.solveMod(9) || eq.solveMod(16) || eq.solveMod(25) {
		sol.codepath = "nomod"
		return
	}

	if eq.IsHyperbolic() {
		return eq.solveHyperbolic()
	}
	if eq.IsElliptical() {
		return eq.solveElliptical()
	}
	if eq.IsParabolic() {
		return eq.solveParabolic()
	}
	panic("not implemented")
}
