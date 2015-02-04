package natural

import (
	"reflect"
	"testing"
)

func TestFactors(t *testing.T) {

	fact := Factorize(13195)
	exp := map[int64]int{5: 1, 7: 1, 13: 1, 29: 1}
	if !reflect.DeepEqual(fact, exp) {
		t.Error("invalid factors", fact)
	}

	fact = Factorize(14)
	exp = map[int64]int{2: 1, 7: 1}
	if !reflect.DeepEqual(fact, exp) {
		t.Error("invalid factors", fact)
	}

	fact = Factorize(48)
	exp = map[int64]int{2: 4, 3: 1}
	if !reflect.DeepEqual(fact, exp) {
		t.Error("invalid factors", fact)
	}
}

func TestNumDivisors(t *testing.T) {

	if NumDivisors(28) != 6 {
		t.Error(NumDivisors(28))
	}
}
