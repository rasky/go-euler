package natural

import (
	"reflect"
	"testing"
)

func TestPermutations(t *testing.T) {

	data := []int64{0, 1, 2}
	exp := [][]int64{
		[]int64{0, 2, 1},
		[]int64{1, 0, 2},
		[]int64{1, 2, 0},
		[]int64{2, 0, 1},
		[]int64{2, 1, 0},
	}

	idx := 0
	for NextPermutation(data) {
		t.Log(data, exp[idx])
		if !reflect.DeepEqual(data, exp[idx]) {
			t.Error("no match")
		}
		idx++
	}

	if idx != len(exp) {
		t.Error("not finished")
	}
}

func TestCombinations(t *testing.T) {

	test := func(data []int64, exp [][]int64) {
		idx := 0
		for NextCombination(data, 3) {
			t.Log(data[:3], exp[idx])
			if !reflect.DeepEqual(data[:3], exp[idx]) {
				t.Error("no match")
			}
			idx++
		}

		if idx != len(exp) {
			t.Error("not finished")
		}

	}

	test([]int64{0, 2, 4, 6, 8},
		[][]int64{
			[]int64{0, 2, 6},
			[]int64{0, 2, 8},
			[]int64{0, 4, 6},
			[]int64{0, 4, 8},
			[]int64{0, 6, 8},
			[]int64{2, 4, 6},
			[]int64{2, 4, 8},
			[]int64{2, 6, 8},
			[]int64{4, 6, 8},
		})

	test([]int64{0, 0, 0, 1, 1, 2},
		[][]int64{
			[]int64{0, 0, 1},
			[]int64{0, 0, 2},
			[]int64{0, 1, 1},
			[]int64{0, 1, 2},
			[]int64{1, 1, 2},
		})

}
