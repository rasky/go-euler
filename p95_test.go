package euler

import "testing"

func TestP95(t *testing.T) {

	sumdivisors := make([]int64, 1000000)
	for i := 1; i < len(sumdivisors)/2; i++ {
		for j := i * 2; j < len(sumdivisors); j += i {
			sumdivisors[j] += int64(i)
		}
	}

	amicablechain := func(n int64) (int, int64) {
		seen := make(map[int64]bool, 16)
		first := n
		cnt := 0
		min := n
		for {
			n = sumdivisors[n]
			cnt += 1
			if n > 1000000 {
				break
			}
			if min > n {
				min = n
			}
			if n == first {
				return cnt, min
			}
			if _, found := seen[n]; found {
				break
			}
			seen[n] = true
		}
		return 0, 0
	}

	tlen, tmin := amicablechain(28)
	if tlen != 1 || tmin != 28 {
		t.Error(tlen, tmin)
	}
	tlen, tmin = amicablechain(284)
	if tlen != 2 || tmin != 220 {
		t.Error(tlen, tmin)
	}
	tlen, tmin = amicablechain(15472)
	if tlen != 5 || tmin != 12496 {
		t.Error(tlen, tmin)
	}

	bestlen, bestmin := 0, int64(0)
	for i := int64(2); i < 1000000; i++ {
		len, min := amicablechain(i)
		if len == 0 {
			continue
		}
		t.Log("Found chain:", len, min)
		if bestlen < len {
			bestlen = len
			bestmin = min
			t.Log("---> New best!")
		} else if bestlen == len && bestmin > min {
			bestmin = min
			t.Log("---> New best!")
		}
	}

	if bestmin != 14316 {
		t.Error(bestmin)
	}
}
