package diophantine

import (
	"fmt"
	"math/big"
	"strings"
)

const (
	SOLUTION_NONE     = 0
	SOLUTION_INFINITE = -1
	SOLUTION_ALL      = -2
)

type CoeffEq2 struct {
	X, Y [3]*big.Int
}

// A structure describing the solution(s) of a diophantine equations
type Solution struct {

	// Number of solutions for the equation. It can be a positive number, or
	// any SOLUTION_* constant.
	NumSolutions int

	// Channel to iteratively generate all solutions; if the number of solutions
	// is finished, the channel will be closed.
	Solutions chan Coords

	// Internal coefficients used for calculating the solution
	Coeff []CoeffEq2

	// Channel used to force-quite the goroutine
	quit chan bool

	// Internal codepath used to solve the equation -- used for debug
	codepath string
}

func (c CoeffEq2) Equal(c2 CoeffEq2) bool {
	return c.X[0].Cmp(c2.X[0]) == 0 && c.X[1].Cmp(c2.X[1]) == 0 && c.X[2].Cmp(c2.X[2]) == 0 &&
		c.Y[0].Cmp(c2.Y[0]) == 0 && c.Y[1].Cmp(c2.Y[1]) == 0 && c.Y[2].Cmp(c2.Y[2]) == 0
}

func (c CoeffEq2) String() string {
	xterms := make([]string, 0)
	yterms := make([]string, 0)
	if c.X[2].Cmp(bzero) != 0 {
		xterms = append(xterms, fmt.Sprintf("%su²", formatTermBig(c.X[2])))
	}
	if c.Y[2].Cmp(bzero) != 0 {
		yterms = append(yterms, fmt.Sprintf("%su²", formatTermBig(c.Y[2])))
	}
	if c.X[1].Cmp(bzero) != 0 {
		xterms = append(xterms, fmt.Sprintf("%su", formatTermBig(c.X[1])))
	}
	if c.Y[1].Cmp(bzero) != 0 {
		yterms = append(yterms, fmt.Sprintf("%su", formatTermBig(c.Y[1])))
	}
	if c.X[0].Cmp(bzero) != 0 {
		xterms = append(xterms, fmt.Sprintf("%+d", c.X[0]))
	}
	if c.Y[0].Cmp(bzero) != 0 {
		yterms = append(yterms, fmt.Sprintf("%+d", c.Y[0]))
	}

	xeq := strings.Join(xterms, "")
	yeq := strings.Join(yterms, "")
	if xeq[0] == '+' {
		xeq = xeq[1:]
	} else if xeq == "" {
		xeq = "0"
	}
	if yeq[0] == '+' {
		yeq = yeq[1:]
	} else if yeq == "" {
		yeq = "0"
	}

	return "[x=" + xeq + ", y=" + yeq + "]"
}

func (sol Solution) String() string {
	switch sol.NumSolutions {
	case SOLUTION_INFINITE:
		if len(sol.Coeff) == 1 {
			return fmt.Sprintf("Solution(INFINITE, %s, %v)", sol.codepath, sol.Coeff[0])
		}
		return fmt.Sprintf("Solution(INFINITE, %s, %v)", sol.codepath, sol.Coeff)
	case SOLUTION_NONE:
		return fmt.Sprintf("Solution(NONE, %s)", sol.codepath)
	case SOLUTION_ALL:
		return fmt.Sprintf("Solution(ALL, %s)", sol.codepath)
	}
	return "Solution(???)"
}

// Finish processing the solution. This function is usually called in a defer
// statement to close the object correctly after it is not needed anymore.
func (sol Solution) Finish() {
	if sol.quit != nil {
		sol.quit <- true
	}
}

type solGenType int

const (
	GEN_ALL_INTEGERS solGenType = iota
	GEN_ALL_INTEGERS_NONZERO
)

func (sol *Solution) generate(gentype solGenType, gen func(n int64) (*big.Int, *big.Int)) {

	sol.NumSolutions = SOLUTION_INFINITE
	sol.Solutions = make(chan Coords)
	sol.quit = make(chan bool, 1)

	gen_one := func(n int64) bool {
		x, y := gen(n)
		if x == nil {
			return false
		}
		s := Coords{X: x, Y: y}
		select {
		case sol.Solutions <- s:
			return true
		case <-sol.quit:
			return false
		}
	}

	switch gentype {
	case GEN_ALL_INTEGERS:
		go func() {
			defer close(sol.Solutions)
			for i := int64(0); ; i++ {
				n := (i + 1) / 2
				if i%2 == 0 {
					n = -n
				}
				if !gen_one(n) {
					return
				}
			}
		}()
	case GEN_ALL_INTEGERS_NONZERO:
		go func() {
			defer close(sol.Solutions)
			for i := int64(1); ; i++ {
				n := (i + 1) / 2
				if i%2 == 0 {
					n = -n
				}
				if !gen_one(n) {
					return
				}
			}
		}()
	}
}
