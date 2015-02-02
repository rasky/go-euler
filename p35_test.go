package euler

import (
	"euler/natural"
	"testing"
)

func TestP35(t *testing.T) {

	primes := make(map[int32]interface{})
	for p := range natural.Primes(1000000) {
		primes[p] = nil
	}

	count := 0
	for p, _ := range primes {
		dig := natural.Digits(int(p))
		dig2 := make([]byte, len(dig))
		circular := true
		for i := 1; i < len(dig); i++ {
			copy(dig2[:len(dig)-i], dig[i:])
			copy(dig2[len(dig)-i:], dig[:i])
			p2 := natural.FromDigits(dig2)
			if _, found := primes[int32(p2)]; !found {
				circular = false
				break
			}
		}
		if circular {
			count += 1
		}
	}

	if count != 55 {
		t.Error(count)
	}
}
