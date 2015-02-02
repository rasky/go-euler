package euler

import (
	"euler/natural"
	"math/big"
	"testing"
)

func TestP53(t *testing.T) {

	cutoff := big.NewInt(1000000)
	count := 0
	for row := range natural.PascalTriangle(23, 101) {
		t.Log(row)
		for _, v := range row {
			if v.Cmp(cutoff) > 0 {
				count += 1
			}
		}
	}

	if count != 4075 {
		t.Error(count)
	}
}
