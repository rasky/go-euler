package euler

import (
	"go-euler/natural"
	"testing"
)

func TestP85(t *testing.T) {

	best := int64(99999)
	bestarea := 0

	i1 := 1
	for t1 := range natural.Triangular(200000) {
		var t2x int64
		i2 := 1
		for t2 := range natural.Triangular(2000000) {
			numrect := t2 * t1
			if numrect > 2000000 {
				if numrect-2000000 < best {
					best = numrect - 2000000
					bestarea = i1 * i2
					t.Log(i1, i2, numrect)
				}
				numrect = t2x * t1
				if 2000000-numrect < best {
					best = 2000000 - numrect
					bestarea = i1 * (i2 - 1)
					t.Log(i1, i2, numrect)
				}
				break
			}
			t2x = t2
			i2++
		}
		i1++
	}

	if bestarea != 2772 {
		t.Error(bestarea)
	}
}
