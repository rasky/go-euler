package natural

import (
	"sort"

	"code.google.com/p/intmath/intgr"
)

var divMemo map[int64][]int

func init() {
	divMemo = make(map[int64][]int)
}

func NumDivisors(x int64) int {

	factors := Factorize(x)

	num := 1
	for _, v := range factors {
		num *= int(v) + 1
	}
	return num
}

func genDivisors(x int64, p []int, rep []int) []int {
	if x == 1 {
		return []int{1}
	}
	if div, found := divMemo[x]; found {
		return div
	}

	n := len(p) - 1
	if n != len(rep)-1 {
		panic("invalid arguments")
	}

	xq := x / int64(intgr.Pow(p[n], rep[n]))

	div := genDivisors(xq, p[:n], rep[:n])
	ret := make([]int, len(div)*(rep[n]+1))
	copy(ret, div)
	offset := len(div)
	for i := 1; i <= rep[n]; i++ {
		for j := offset * i; j < offset*(i+1); j++ {
			ret[j] = ret[j-offset] * p[n]
		}
	}

	sort.Ints(ret)
	divMemo[x] = ret

	return ret
}

func Divisors(x int64) []int {
	if x < 2 {
		return []int{}
	}

	factors := Factorize(x)

	p := make([]int, 0, len(factors))
	rep := make([]int, 0, len(factors))

	for f := range factors {
		p = append(p, int(f))
	}
	sort.Ints(p)
	for _, f := range p {
		rep = append(rep, int(factors[int64(f)]))
	}

	div := genDivisors(x, p, rep)
	return div[0 : len(div)-1]
}
