package natural

import (
	"reflect"
	"testing"
)

func allFactors(x int64) []int64 {
	fact := make([]int64, 0, 8)
	for f := range Factors(x) {
		fact = append(fact, f)
	}
	return fact
}

func TestFactors(t *testing.T) {

	fact := allFactors(13195)
	exp := []int64{5, 7, 13, 29}
	if !reflect.DeepEqual(fact, exp) {
		t.Error("invalid factors", fact)
	}

	fact = allFactors(14)
	exp = []int64{2, 7}
	if !reflect.DeepEqual(fact, exp) {
		t.Error("invalid factors", fact)
	}

	fact = allFactors(48)
	exp = []int64{2, 2, 2, 2, 3}
	if !reflect.DeepEqual(fact, exp) {
		t.Error("invalid factors", fact)
	}
}

func TestNumDivisors(t *testing.T) {

	if NumDivisors(28) != 6 {
		t.Error(NumDivisors(28))
	}
}
