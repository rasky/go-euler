package euler

import (
	"euler/natural"
	"testing"
)

func TestP37(t *testing.T) {

	primes := make(map[int32]interface{})
	for p := range natural.Primes(1000000) {
		primes[p] = nil
	}

	sum := 0
	count := 0
	for p, _ := range primes {
		if p < 10 {
			continue
		}

		dig := natural.Digits(int(p))

		truncatable := true
		for i := 1; i < len(dig); i++ {
			p2 := natural.FromDigits(dig[i:])
			if _, found := primes[int32(p2)]; !found {
				truncatable = false
				break
			}
			p3 := natural.FromDigits(dig[:len(dig)-i])
			if _, found := primes[int32(p3)]; !found {
				truncatable = false
				break
			}
		}
		if truncatable {
			sum += int(p)
			count += 1
			if count == 11 {
				break
			}
		}
	}

	if count != 11 {
		t.Error(count)
	}

	if sum != 748317 {
		t.Error(sum)
	}

}
