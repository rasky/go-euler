package euler

import (
	"go-euler/natural"
	"testing"
)

func TestP24(t *testing.T) {
	p := natural.IthPermutation(10, 1000000-1)
	total := int64(0)
	for _, dig := range p {
		total *= 10
		total += dig
	}
	if total != 2783915460 {
		t.Error(total, p)
	}
}
