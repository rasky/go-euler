package euler

import "testing"

func TestP47(t *testing.T) {

	consecFactors := func(search int) int {
		sieve := make([]int, 1000000)
		count := 0
		for i := 2; i < len(sieve); i++ {
			if sieve[i] == 0 {
				count = 0
				for j := i + i; j < len(sieve); j += i {
					sieve[j]++
				}
			} else if sieve[i] == search {
				count++
				if count == search {
					return i - search + 1
				}
			} else {
				count = 0
			}
		}
		return -1
	}

	c2 := consecFactors(2)
	c3 := consecFactors(3)
	c4 := consecFactors(4)

	if c2 != 14 || c3 != 644 || c4 != 134043 {
		t.Error(c2, c3, c4)
	}
}
