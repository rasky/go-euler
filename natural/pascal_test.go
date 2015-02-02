package natural

import (
	"math/big"
	"reflect"
	"testing"
)

func TestPascalTriangle(t *testing.T) {

	matchRow := func(exp []int64, row []*big.Int) bool {
		row2 := make([]int64, len(row))
		for idx, b := range row {
			row2[idx] = b.Int64()
		}
		return reflect.DeepEqual(row2, exp)
	}

	ch := PascalTriangle(0, 3)
	r := <-ch
	if !matchRow([]int64{1}, r) {
		t.Error(r)
	}
	r = <-ch
	if !matchRow([]int64{1, 1}, r) {
		t.Error(r)
	}
	r = <-ch
	if !matchRow([]int64{1, 2, 1}, r) {
		t.Error(r)
	}
	r, ok := <-ch
	if ok {
		t.Error("channel not closed")
	}

	ch = PascalTriangle(23, 26)
	row23 := []int64{1, 23, 253, 1771, 8855, 33649, 100947, 245157, 490314, 817190, 1144066, 1352078, 1352078, 1144066, 817190, 490314, 245157, 100947, 33649, 8855, 1771, 253, 23, 1}
	row24 := []int64{1, 24, 276, 2024, 10626, 42504, 134596, 346104, 735471, 1307504, 1961256, 2496144, 2704156, 2496144, 1961256, 1307504, 735471, 346104, 134596, 42504, 10626, 2024, 276, 24, 1}
	row25 := []int64{1, 25, 300, 2300, 12650, 53130, 177100, 480700, 1081575, 2042975, 3268760, 4457400, 5200300, 5200300, 4457400, 3268760, 2042975, 1081575, 480700, 177100, 53130, 12650, 2300, 300, 25, 1}

	r = <-ch
	if !matchRow(row23, r) {
		t.Error(r)
	}
	r = <-ch
	if !matchRow(row24, r) {
		t.Error(r)
	}
	r = <-ch
	if !matchRow(row25, r) {
		t.Error(r)
	}
	r, ok = <-ch
	if ok {
		t.Error("channel not closed")
	}
}
