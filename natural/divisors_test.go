package natural

import (
	"reflect"
	"testing"
)

func TestDivisors(t *testing.T) {

	div := Divisors(220)
	exp := []int{1, 2, 4, 5, 10, 11, 20, 22, 44, 55, 110}
	if !reflect.DeepEqual(div, exp) {
		t.Error(div)
	}

	div = Divisors(284)
	exp = []int{1, 2, 4, 71, 142}
	if !reflect.DeepEqual(div, exp) {
		t.Error(div)
	}
}
