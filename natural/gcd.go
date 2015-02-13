package natural

import "code.google.com/p/intmath/i64"

func GCD(a, b int64) int64 {
	a = i64.Abs(a)
	b = i64.Abs(b)
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func LCM(a, b int64) int64 {
	return a * b / GCD(a, b)
}
