package euler

import (
	"go-euler/natural"
	"strings"
	"testing"
)

func TestP79(t *testing.T) {

	var used [10]bool
	seen := make(map[string]bool)
	attempts := make([]string, 0, 100)
	for _, line := range strings.Split(p79attempts[1:], "\n") {
		if _, found := seen[line]; found {
			continue

		}
		for _, a := range line {
			used[a-'0'] = true
		}
		attempts = append(attempts, line)
	}

	check := func(dig []byte, attempt string) bool {
		idx := 0
		for _, d := range dig {
			if byte(attempt[idx]-'0') == d {
				idx++
				if idx == len(attempt) {
					return true
				}
			}
		}
		return false
	}

	// NOTE: this solution doesn't handle repetitions of digits
	dig := make([]byte, 0, 10)
	for d, found := range used {
		if found {
			dig = append(dig, byte(d))
		}
	}

	idx := 0
	for natural.DigitsNextPermutation(dig) {
		error := false
		for _, a := range attempts {
			if !check(dig, a) {
				error = true
				break
			}
		}
		if !error {
			break
		}
		idx++
	}

	if natural.FromDigitsReverse(dig) != 73162890 {
		t.Error(natural.FromDigitsReverse(dig))
	}
}

var p79attempts string = `
319
680
180
690
129
620
762
689
762
318
368
710
720
710
629
168
160
689
716
731
736
729
316
729
729
710
769
290
719
680
318
389
162
289
162
718
729
319
790
680
890
362
319
760
316
729
380
319
728
716`
