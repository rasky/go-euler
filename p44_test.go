package euler

import (
	"euler/natural"
	"testing"
)

func TestP44(t *testing.T) {

	pents := make([]int64, 0, 1000)
	ch := natural.Pentagonal(10000000)

	pents = append(pents, <-ch)
	pents = append(pents, <-ch)
	a, b, c := 0, 0, 0

Loop:
	for {
		a = 0
		b = len(pents) - 1
		next := <-ch
		diff := next - pents[b]
		for a < b-2 && pents[a] <= diff {
			a++
		}
		for c = a; c < b; c++ {
			if natural.IsPentagonal(pents[c]+pents[b]) && natural.IsPentagonal(pents[b]-pents[c]) {
				break Loop
			}
		}

		pents = append(pents[a:], next)
	}

	if pents[b]-pents[c] != 5482660 {
		t.Error(pents[b] - pents[c])
	}
}
