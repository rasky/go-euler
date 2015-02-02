package natural

import (
	"fmt"
	"math/big"

	"code.google.com/p/intmath/i64"
)

type Rational struct {
	Num, Den int64
}

func GCD(a, b int64) int64 {
	a = i64.Abs(a)
	b = i64.Abs(b)
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func LCM(a, b int64) int64 {
	return a * b / GCD(a, b)
}

func (r1 Rational) Reciprocal() Rational {
	return Rational{r1.Den, r1.Num}
}

func (r1 Rational) Equal(r2 Rational) bool {
	lcm := LCM(r1.Den, r2.Den)
	return r1.Num*(lcm/r1.Den) == r2.Num*(lcm/r2.Den)
}

func (r1 Rational) Sum(r2 Rational) Rational {
	lcm := LCM(r1.Den, r2.Den)
	num1 := r1.Num * (lcm / r1.Den)
	num2 := r2.Num * (lcm / r2.Den)
	return Rational{num1 + num2, lcm}
}

func (r1 Rational) Reduce() Rational {

	for f := range Factors(int64(r1.Den)) {
		if r1.Num%f == 0 {
			r1.Num /= f
			r1.Den /= f
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

func ContinuedFraction(terms []int64) *big.Rat {

	n := len(terms)
	r := big.NewRat(terms[n-1], 1)
	for i := n - 2; i >= 0; i-- {
		r.Inv(r)
		r.Add(r, big.NewRat(terms[i], 1))
	}
	return r
}
