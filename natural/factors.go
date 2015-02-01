package natural

import (
	"sort"

	"code.google.com/p/intmath/i64"
)

var factMemo map[int64]map[int32]int32

func init() {
	factMemo = make(map[int64]map[int32]int32)
}

func Factors(x int64) chan int32 {

	ch := make(chan int32)

	go func() {
		defer close(ch)

		xx := x
		memo := make(map[int32]int32, 32)
		defer func() { factMemo[xx] = memo }()

		replay := func(x int64) {

			fk := make([]int, len(factMemo[x]))
			i := 0
			for k, _ := range factMemo[x] {
				fk[i] = int(k)
				i += 1
			}
			sort.Ints(fk)

			for _, k := range fk {
				num := factMemo[x][int32(k)]
				for i := int32(0); i < num; i++ {
					ch <- int32(k)
				}
				memo[int32(k)] += num
			}
		}

		for x&1 == 0 {
			x = x >> 1
			ch <- 2
			memo[2] += 1
		}
		limit := i64.Sqrt(x)
		for y := range Primes(int32(limit)) {
			if _, found := factMemo[x]; found {
				replay(x)
				return
			}
			for (x % int64(y)) == 0 {
				x = x / int64(y)
				ch <- y
				memo[y] += 1
				if _, found := factMemo[x]; found {
					replay(x)
					return
				}
			}
		}
		if x != 1 {
			ch <- int32(x)
			memo[int32(x)] += 1
		}

		factMemo[xx] = memo
	}()

	return ch
}

func Factorize(x int64) map[int32]int32 {

	factors := make(map[int32]int32, 32)
	for f := range Factors(x) {
		factors[f] += 1
	}

	return factors
}
