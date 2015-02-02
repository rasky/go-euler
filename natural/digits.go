package natural

func DigitsWithBase(x, base int) []byte {
	digits := make([]byte, 0, 8)
	for x != 0 {
		digits = append(digits, byte(x%base))
		x /= base
	}
	return digits
}

func Digits(x int) []byte {
	return DigitsWithBase(x, 10)
}

func FromDigits(dig []byte) int {
	num := 0
	for i := len(dig) - 1; i >= 0; i-- {
		num *= 10
		num += int(dig[i])
	}
	return num
}

func FromDigitsReverse(dig []byte) int {
	dig2 := make([]byte, len(dig))
	copy(dig2, dig)
	Reverse(dig2)
	return FromDigits(dig2)
}

// Concatenate numbers (eg: Concat({12,34}) == 1234)
func Concat(x []int) int {

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
