package euler

import (
	"go-euler/natural"
	"strconv"
	"testing"
)

func TestP68(t *testing.T) {

	mgon := func(n int64) []string {

		solutions := make([]string, 0, 8)

		data := make([]int64, n)
		itoa := make([]string, n+1)
		for i := int64(0); i < n; i++ {
			data[i] = i + 1
			itoa[i+1] = strconv.Itoa(int(i + 1))
		}

		for {
			valid := true

			// Check order
			min := data[0]
			for i := int64(3); i < n; i += 2 {
				if data[i] < min {
					valid = false
					break
				}
			}
			if valid {
				// Check sum
				sum := data[0] + data[1] + data[2]
				for i := int64(3); i < n; i += 2 {
					j := i + 1
					if j == n {
						j = 1
					}
					if sum != data[i]+data[i-1]+data[j] {
						valid = false
						break
					}
				}
			}
			if valid {
				s := itoa[data[0]] + itoa[data[1]] + itoa[data[2]]
				for i := int64(3); i < n; i += 2 {
					j := i + 1
					if j == n {
						j = 1
					}
					s += itoa[data[i]] + itoa[data[i-1]] + itoa[data[j]]
				}
				solutions = append(solutions, s)
			}

			if !natural.NextPermutation(data) {
				break
			}
		}
		return solutions
	}

	solutions := mgon(6)
	best := ""
	for _, s := range solutions {
		if best < s {
			best = s
		}
	}

	if best != "432621513" {
		t.Error(best)
	}

	solutions = mgon(10)
	best = ""
	for _, s := range solutions {
		if len(s) == 16 && best < s {
			best = s
		}
	}

	if best != "6531031914842725" {
		t.Error(best)
	}

}
