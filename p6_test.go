package euler

import "testing"

func p6(n int) int {
	total := 0
	for i := 1; i <= n; i++ {
		total += i * i
	}

	sq := (n * (n + 1)) / 2

	return sq*sq - total
}

func TestP6(t *testing.T) {

	p6a := p6(10)
	if p6a != 2640 {
		t.Error(p6a)
	}

	p6b := p6(100)
	if p6b != 25164150 {
		t.Error(p6b)
	}

}
