package euler

import (
	"euler/natural"
	"testing"
)

func p23abundant(x int64) bool {
	sum := int64(0)
	for _, d := range natural.Divisors(x) {
		sum += int64(d)
	}
	return (sum > x)
}

func TestP23(t *testing.T) {

	if !p23abundant(12) {
		t.Error("12 should be abundant")
	}

	ba := natural.NewBitArray(60000)

	abund := make([]int64, 0, 100)
	for i := int64(12); i <= 28123; i++ {
		if p23abundant(i) {
			abund = append(abund, i)
			for _, a := range abund {
				ba.Set(a + i)
			}
		}
	}

	sum := int64(0)
	for i := int64(1); i <= 28123; i++ {
		if !ba.Get(i) {
			sum += i
		}
	}

	if sum != 4179871 {
		t.Error(sum)
	}
}
