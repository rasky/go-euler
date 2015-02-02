package natural

import "reflect"

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
	Reverse(dig2)
	return FromDigits(dig2)
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
func Reverse(x []byte) {
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
	Reverse(x[i+1:])
	return true
}

func IsPermutation(x1 []byte, x2 []byte) bool {

	m1 := make(map[byte]int, 10)
	m2 := make(map[byte]int, 10)

	for _, x := range x1 {
		m1[x]++
	}
	for _, x := range x2 {
		m2[x]++
	}

	return reflect.DeepEqual(m1, m2)
}
