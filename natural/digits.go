package natural

import (
	"reflect"

	"code.google.com/p/intmath/intgr"
)

func DigitsWithBase(x int64, base int) []byte {
	digits := make([]byte, 0, 8)
	for x != 0 {
		digits = append(digits, byte(x%int64(base)))
		x /= int64(base)
	}
	return digits
}

func Digits(x int64) []byte {
	return DigitsWithBase(x, 10)
}

func FromDigits(dig []byte) int64 {
	num := int64(0)
	for i := len(dig) - 1; i >= 0; i-- {
		num *= 10
		num += int64(dig[i])
	}
	return num
}

func FromDigitsReverse(dig []byte) int64 {
	dig2 := make([]byte, len(dig))
	copy(dig2, dig)
	DigitsReverse(dig2)
	return FromDigits(dig2)
}

// Compute x+y in digit form
func DigitsAdd(x []byte, y []byte) []byte {

	n := intgr.Max(len(x), len(y))
	res := make([]byte, n+1)

	c := byte(0)
	for i := 0; i < n; i++ {
		xx, yy := byte(0), byte(0)
		if i < len(x) {
			xx = x[i]
		}
		if i < len(y) {
			yy = y[i]
		}
		v := xx + yy + c
		res[i] = v % 10
		c = v / 10
	}
	res[n] = c
	if c == 0 {
		return res[:n]
	}
	return res
}

func DigitsPalindromic(dig []byte) bool {
	for i := 0; i < len(dig)/2; i++ {
		if dig[i] != dig[len(dig)-i-1] {
			return false
		}
	}
	return true
}

// Concatenate numbers (eg: Concat({12,34}) == 1234)
func Concat(x []int64) int64 {

	dig := make([]byte, 0, 16)
	for i := len(x) - 1; i >= 0; i-- {
		dig = append(dig, Digits(x[i])...)
	}

	return FromDigits(dig)
}

// Reverse digits
func DigitsReverse(x []byte) {
	l := len(x)
	for i := 0; i < l/2; i++ {
		x[i], x[l-i-1] = x[l-i-1], x[i]
	}
}

func NextPermutation(x []byte) bool {
	var i, j int
	for i = len(x) - 2; i >= 0; i-- {
		if x[i] < x[i+1] {
			break
		}
	}
	if i < 0 {
		return false
	}

	for j = len(x) - 1; j > i+1; j-- {
		if x[j] > x[i] {
			break
		}
	}

	x[i], x[j] = x[j], x[i]
	DigitsReverse(x[i+1:])
	return true
}

func IsPermutationMulti(x1 []byte, x2 [][]byte) bool {

	var m1 [10]int
	for _, x := range x1 {
		m1[x]++
	}
	for _, xx := range x2 {
		if len(xx) != len(x1) {
			return false
		}
		var m2 [10]int
		for _, x := range xx {
			m2[x]++
		}
		if !reflect.DeepEqual(m1, m2) {
			return false
		}
	}
	return true
}

func IsPermutation(x1 []byte, x2 []byte) bool {
	return IsPermutationMulti(x1, [][]byte{x2})
}
