package euler

import "testing"

func p31CoinSums(total, partial int, coins []int) int {
	if coins[0] == 1 {
		return 1
	}
	if total == partial {
		return 1
	}

	sums := 0
	missing := total - partial
	for i := 0; i <= missing/coins[0]; i++ {
		sums += p31CoinSums(total, partial+i*coins[0], coins[1:])
	}
	return sums
}

func TestP31(t *testing.T) {
	coins := []int{200, 100, 50, 20, 10, 5, 2, 1}

	sums := p31CoinSums(200, 0, coins)
	if sums != 73682 {
		t.Error(sums)
	}
}
