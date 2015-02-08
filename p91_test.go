package euler

import (
	"go-euler/natural"
	"testing"
)

func TestP91(t *testing.T) {

	solve := func(n int64) int64 {
		count := n * n * 3

		for x := int64(1); x <= n; x++ {
			for y := int64(1); y <= n; y++ {
				slope := natural.GCD(x, y)
				tri1 := y * slope / x
				tri2 := (n - x) * slope / y
				count += natural.Min(tri1, tri2) * 2
			}
		}

		return count
	}

	s6 := solve(2)
	if s6 != 14 {
		t.Error(s6)
	}

	s50 := solve(50)
	if s50 != 14234 {
		t.Error(s50)
	}
}
