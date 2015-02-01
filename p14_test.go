package euler

import (
	"euler/natural"
	"testing"
)

func TestP14(t *testing.T) {

	var max int
	var maxv int64

	for i := int64(1); i < int64(1000000); i++ {
		l := natural.LenCollatz(i)
		if max < l {
			max = l
			maxv = i
		}
	}

	if maxv != 837799 {
		t.Error(maxv, max)
	}
}
