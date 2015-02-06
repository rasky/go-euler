package euler

import (
	"go-euler/natural"
	"testing"
)

func TestP87(t *testing.T) {

	const TARGET = int64(50000000)
	limq, limt, lims := 0, 0, 0

	primes := make([]int64, 0, 1024)
	for p := range natural.Primes(10000) {
		primes = append(primes, p)
		p2 := p * p
		if p2 < TARGET {
			lims = len(primes)
		}
		p3 := p2 * p
		if p3 < TARGET {
			limt = len(primes)
		}
		p4 := p3 * p
		if p4 < TARGET {
			limq = len(primes)
		}
	}
	t.Log(primes[lims], primes[limt], primes[limq])
	t.Log(primes[lims-1], primes[limt-1], primes[limq-1])

	seen := make(map[int64]bool, 2000000)
	for i := 0; i < limq; i++ {
		pi := primes[i] * primes[i] * primes[i] * primes[i]
		for j := 0; j < limt; j++ {
			pj := pi + primes[j]*primes[j]*primes[j]
			if pj >= TARGET {
				break
			}
			var z int
			for z = 0; z < lims; z++ {
				pz := pj + primes[z]*primes[z]
				if pz >= TARGET {
					break
				}
				seen[pz] = true
			}
		}
	}

	if len(seen) != 1097343 {
		t.Error(len(seen))
	}

}
