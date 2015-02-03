package euler

import (
	"go-euler/natural"
	"testing"
)

func TestP55(t *testing.T) {

	count := 0
	for i := int64(1); i < 10000; i++ {
		if natural.IsLychrel(i, 50) {
			count++
		}
	}

	if count != 249 {
		t.Error(count)
	}
}
