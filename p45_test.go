package euler

import (
	"go-euler/natural"
	"testing"
)

func TestP45(t *testing.T) {

	match := make([]int64, 0)

	for i := range natural.Hexagonal(10000000000) {
		if natural.IsTriangular(i) && natural.IsPentagonal(i) {
			match = append(match, i)
			if len(match) == 3 {
				break
			}
		}
	}

	if match[0] != 1 || match[1] != 40755 || match[2] != 1533776805 {
		t.Error(match)
	}
}
