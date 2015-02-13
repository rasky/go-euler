package diophantine

import (
	"fmt"
	"math/big"
	"testing"
)

func TestContFractSqrt(t *testing.T) {

	cf := NewContinuedFractionForSqrt(114, 0, 1)
	exp := ContinuedFraction{
		FixedTerms:  []int64{10},
		PeriodTerms: []int64{1, 2, 10, 2, 1, 20},
	}
	if !cf.Equal(exp) {
		t.Error(cf)
	}

	cf = NewContinuedFractionForSqrt(313, -41, 38)
	exp = ContinuedFraction{
		FixedTerms:  []int64{-1, 2},
		PeriodTerms: []int64{1, 1, 2, 2, 1, 1, 3, 1, 5, 8, 1, 2, 17, 2, 1, 8, 5, 1, 3},
	}
	if !cf.Equal(exp) {
		t.Error(cf)
	}

	cf = NewContinuedFractionForSqrt(-313, -41, 38)
	exp = ContinuedFraction{
		FixedTerms:  []int64{-2, 2},
		PeriodTerms: []int64{1, 1, 2, 2, 1, 1, 3, 1, 5, 8, 1, 2, 17, 2, 1, 8, 5, 1, 3},
	}
	if !cf.Equal(exp) {
		t.Error(cf)
	}
}

func TestContFractEval(t *testing.T) {

	cf := ContinuedFraction{
		FixedTerms:  []int64{-2, 1, 5},
		PeriodTerms: []int64{8, 5, 1, 3, 1, 1, 2, 2, 1, 1, 3, 1, 5, 8, 1, 2, 17, 2, 1},
	}

	rat := cf.Evaluate(40)
	var num, den big.Int
	fmt.Sscan("-423629283087437383", &num)
	fmt.Sscan("364150459879840134", &den)
	if rat.Num().Cmp(&num) != 0 || rat.Denom().Cmp(&den) != 0 {
		t.Error(rat)
	}
}
