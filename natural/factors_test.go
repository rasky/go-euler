package natural

import (
	"reflect"
	"testing"
)

func allFactors(x int64) []int32 {
	fact := make([]int32, 0, 8)
	for f := range Factors(x) {
		fact = append(fact, f)
	}
	return fact
}

func TestFactors(t *testing.T) {

	fact := allFactors(13195)
	exp := []int32{5, 7, 13, 29}
	if !reflect.DeepEqual(fact, exp) {
		t.Error("invalid factors", fact)
	}

	fact = allFactors(14)
	exp = []int32{2, 7}
	if !reflect.DeepEqual(fact, exp) {
		t.Error("invalid factors", fact)
	}

	fact = allFactors(48)
	exp = []int32{2, 2, 2, 2, 3}
	if !reflect.DeepEqual(fact, exp) {
		t.Error("invalid factors", fact)
	}
}

func TestNumDivisors(t *testing.T) {

	if NumDivisors(28) != 6 {
		t.Error(NumDivisors(28))
	}
}
