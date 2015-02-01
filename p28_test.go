package euler

import "testing"

func p28SpiralSum(sz int) int {
	sum := 1
	count := 1
	for i := 3; i <= sz; i += 2 {
		count += (i - 1)
		sum += count
		count += (i - 1)
		sum += count
		count += (i - 1)
		sum += count
		count += (i - 1)
		sum += count
	}
	return sum
}

func TestP28(t *testing.T) {

	sum := p28SpiralSum(5)
	if sum != 101 {
		t.Error(sum)
	}

	sum = p28SpiralSum(1001)
	if sum != 669171001 {
		t.Error(sum)
	}
}
