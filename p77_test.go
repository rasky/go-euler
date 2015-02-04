package euler

import (
	"go-euler/natural"
	"testing"
)

func TestP77(t *testing.T) {

	const target = 10

	primes := make([]int64, 0, 1000)
	for p := range natural.Primes(10000) {
		primes = append(primes, p)
	}

	var countsums func(total int64, primes []int64) int64
	countsums = func(total int64, primes []int64) int64 {
		pidx := len(primes) - 1
		if pidx == 0 {
			return 1 - total%2
		}
		sums := int64(0)
		p := primes[pidx]
		for total > 1 {
			for primes[pidx-1] > total {
				pidx--
			}
			sums += countsums(total, primes[:pidx])
			total -= p
		}
		if total == 0 {
			sums++
		}
		return sums
	}

	var solution int64
	pidx := 0
	for i := int64(10); i < 500; i++ {
		for primes[pidx] < i {
			pidx++
		}
		sums := countsums(i, primes[:pidx])
		t.Log(i, sums)
		if sums > 5000 {
			solution = i
			break
		}
	}

	if solution != 71 {
		t.Error(solution)
	}
}
