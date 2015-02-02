package euler

import (
	"euler/natural"
	"testing"
)

func TestP50(t *testing.T) {

	primes := make([]int64, 0, 1024)

	for p := range natural.Primes(1000000) {
		primes = append(primes, p)
	}

	best := int64(0)
	bestlen := 0

	for i := 0; i < len(primes)-bestlen; i++ {
		total := int64(0)
		for j := i; j < i+bestlen; j++ {
			total += primes[j]
		}
		for j := i + bestlen; j < len(primes); j++ {
			total += primes[j]
			if total >= 1000000 {
				break
			}
			if natural.IsPrime(total) {
				bestlen = j - i + 1
				best = total
				t.Log(best, bestlen, primes[i:i+bestlen])
			}
		}
	}

	if best != 997651 {
		t.Error(best)
	}
}
