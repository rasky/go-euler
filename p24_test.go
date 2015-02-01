package euler

import "testing"

func ith_permutation(n, i int) []int {
	fact := make([]int, n)
	perm := make([]int, n)

	fact[0] = 1
	for k := 1; k < n; k++ {
		fact[k] = fact[k-1] * k
	}

	for k := 0; k < n; k++ {
		perm[k] = i / fact[n-1-k]
		i = i % fact[n-1-k]
	}

	for k := n - 1; k > 0; k-- {
		for j := k - 1; j >= 0; j-- {
			if perm[j] <= perm[k] {
				perm[k]++
			}
		}
	}

	return perm
}

func TestP24(t *testing.T) {
	p := ith_permutation(10, 1000000-1)
	total := 0
	for _, dig := range p {
		total *= 10
		total += dig
	}
	if total != 2783915460 {
		t.Error(total, p)
	}
}
