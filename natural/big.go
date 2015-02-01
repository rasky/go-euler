package natural

import "math/big"

func BigSumDigits(b *big.Int) int {
	sum := 0
	for _, s := range b.String() {
		sum += int(s - '0')
	}
	return sum
}
