package euler

import (
	"go-euler/natural"
	"testing"
)

func TestP39(t *testing.T) {

	max, maxi := 0, 0

	for i := 9; i < 1000; i++ {
		count := 0
		for _ = range natural.PythagoreanTriplesWithSum(i) {
			count++
		}

		if max < count {
			max = count
			maxi = i
		}
	}

	if maxi != 840 {
		t.Error(maxi)
	}
}
