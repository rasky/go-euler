package euler

import (
	"euler/natural"
	"testing"
)

func p23abundant(x int) bool {
	sum := 0
	for _, d := range natural.Divisors(int64(x)) {
		sum += d
	}
	return (sum > x)
}

func TestP23(t *testing.T) {

	if !p23abundant(12) {
		t.Error("12 should be abundant")
	}

	ba := natural.NewBitArray(60000)

	abund := make([]int, 0, 100)
	for i := 12; i <= 28123; i++ {
		if p23abundant(i) {
			abund = append(abund, i)
			for _, a := range abund {
				ba.Set(int32(a + i))
			}
		}
	}

	sum := 0
	for i := 1; i <= 28123; i++ {
		if !ba.Get(int32(i)) {
			sum += i
		}
	}

	if sum != 4179871 {
		t.Error(sum)
	}
}
