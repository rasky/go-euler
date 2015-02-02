package euler

import (
	"euler/natural"
	"testing"

	"code.google.com/p/intmath/i32"
)

func TestP46(t *testing.T) {

	primes := make([]int32, 0)
	n := int32(0)

Loop:
	for p := range natural.Primes(1000000) {
		primes = append(primes, p)
		if len(primes) < 3 {
			continue
		}
		for n = primes[len(primes)-2] + 2; n < p; n += 2 {
			found := false
			for _, m := range primes[:len(primes)-1] {
				q := (n - m) / 2
				sq := i32.Sqrt(q)
				if sq*sq == q {
					found = true
					break
				}
			}
			if !found {
				break Loop
			}
		}
	}

	if n != 5777 {
		t.Error(n)
	}
}
