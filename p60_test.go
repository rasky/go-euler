package euler

import (
	"euler/natural"
	"math"
	"testing"
)

func TestP60(t *testing.T) {

	const MAXPRIME = 60000
	const MEMOW = 8000

	primelist := make([]int64, 0, 1000)
	for p := range natural.Primes(MAXPRIME) {
		primelist = append(primelist, p)
	}
	if MEMOW < len(primelist) {
		panic("memo too small")
	}

	pow10 := []int64{10, 100, 1000, 10000, 100000, 1000000, 10000000}
	concat := func(p1, p2 int64) int64 {
		return p1*pow10[int(math.Log10(float64(p2)))] + p2
	}

	isPrime := func(n int64) bool {
		if n == 2 {
			return true
		}
		for i := 0; ; i++ {
			p := primelist[i]
			if p*p > n {
				return true
			}
			if n%p == 0 {
				return false
			}
		}
	}

	isPairPrime := func(idx1, idx2 int64) bool {
		p1 := primelist[idx1]
		p2 := primelist[idx2]
		return isPrime(concat(p1, p2)) &&
			isPrime(concat(p2, p1))
	}

	var res []int64
	var pairs [MEMOW][MEMOW]bool
Loop:
	for ax := int64(0); ax < int64(len(primelist)); ax++ {
		for bx := int64(0); bx < ax; bx++ {
			if isPairPrime(ax, bx) {
				pairs[bx][ax] = true
			}
		}
		for bx := int64(0); bx < ax; bx++ {
			if !pairs[bx][ax] {
				continue
			}
			for cx := int64(0); cx < bx; cx++ {
				if !pairs[cx][bx] || !pairs[cx][ax] {
					continue
				}
				for dx := int64(0); dx < cx; dx++ {
					if !pairs[dx][cx] || !pairs[dx][bx] || !pairs[dx][ax] {
						continue
					}
					for ex := int64(0); ex < dx; ex++ {
						if !pairs[ex][dx] || !pairs[ex][cx] || !pairs[ex][bx] || !pairs[ex][ax] {
							continue
						}

						res = []int64{primelist[ax], primelist[bx], primelist[cx], primelist[dx], primelist[ex]}
						break Loop
					}
				}
			}
		}
	}

	if natural.Sum(res) != 26033 {
		t.Error(natural.Sum(res))
	}
}

/*

func TestP60(t *testing.T) {

	// This problem is about O(n^5), so we want to keep the N down as much as
	// possible. Since we know that the biggest prime in the solution is 8389,
	// we put 8500 here, but the algorithm works with bigger numbers too
	// (and given enough time to complete...)
	const MAXPRIME = 12000

	primeidx := make(map[int64]int, 1000)
	primelist := make([]int64, 0, 1000)
	for p := range natural.Primes(MAXPRIME) {
		primeidx[p] = len(primelist)
		primelist = append(primelist, p)
	}

	pow10 := []int64{10, 100, 1000, 10000, 100000, 1000000, 10000000}
	concat := func(p1, p2 int64) int64 {
		return p1*pow10[int(math.Log10(float64(p2)))] + p2
	}

	isPairPrime := func(p1, p2 int64) bool {
		return natural.IsPrime(concat(p1, p2)) &&
			natural.IsPrime(concat(p2, p1))
	}

	// For each prime, the sored list of (bigger) paired primes
	pairlist_memo := make(map[int64][]int64)
	pairlist := func(p int64) []int64 {
		if list, found := pairlist_memo[p]; found {
			return list
		}

		list := make([]int64, 0, 128)
		for _, p2 := range primelist[primeidx[p]+1:] {
			if isPairPrime(p, p2) {
				list = append(list, p2)
			}
		}
		pairlist_memo[p] = list
		return list
	}

	intersection := func(l1, l2 []int64) []int64 {
		res := make([]int64, 0, 128)
		for len(l1) > 0 && len(l2) > 0 {
			switch {
			case l1[0] == l2[0]:
				res = append(res, l1[0])
				l1, l2 = l1[1:], l2[1:]
			case l1[0] < l2[0]:
				l1 = l1[1:]
			case l1[0] > l2[0]:
				l2 = l2[1:]
			}
		}
		return res
	}

	var recurse func(group []int64, attempt []int64, best int64) int64
	recurse = func(group []int64, attempts []int64, best int64) int64 {
		for idx, p := range attempts {
			group2 := append(group, p)
			if len(group2) == 5 {
				sum := natural.Sum(group2)
				t.Log(group2, sum)
				if best > sum || best < 0 {
					best = sum
				}
				return best
			}
			attempts2 := intersection(attempts[idx+1:], pairlist(p))
			if len(attempts2) > 0 {
				best = recurse(group2, attempts2, best)
			}
		}
		return best
	}

	best := recurse(make([]int64, 0, 8), primelist[1:], -1)
	if best != 26033 {
		t.Error(best)
	}
}
*/
