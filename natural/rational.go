package natural

import (
	"fmt"

	"code.google.com/p/intmath/intgr"
)

type Rational struct {
	Num, Den int
}

func GCD(a, b int) int {
	a = intgr.Abs(a)
	b = intgr.Abs(b)
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func LCM(a, b int) int {
	return a * b / GCD(a, b)
}

func (r1 Rational) Equal(r2 Rational) bool {
	lcm := LCM(r1.Den, r2.Den)
	return r1.Num*(lcm/r1.Den) == r2.Num*(lcm/r2.Den)
}

func (r1 Rational) Reduce() Rational {

	for f := range Factors(int64(r1.Den)) {
		fi := int(f)
		if r1.Num%fi == 0 {
			r1.Num /= fi
			r1.Den /= fi
		}
	}
	return r1
}

func (r1 Rational) Mul(r2 Rational) Rational {
	r1.Num *= r2.Num
	r1.Den *= r2.Den
	return r1
}

func (r Rational) String() string {
	return fmt.Sprintf("<%d/%d>", r.Num, r.Den)
}
