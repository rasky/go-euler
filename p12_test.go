package euler

import (
	"go-euler/natural"
	"testing"
)

func TestP12(t *testing.T) {

	var num int
	var tr int64
	for tr = range natural.Triangular(1000000000) {
		num = natural.NumDivisors(tr)
		if num > 500 {
			break
		}
	}

	if tr != 76576500 {
		t.Error(tr)
	}
}
