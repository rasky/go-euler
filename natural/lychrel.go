package natural

// Check if n is a Lychrel number, with the specified number of iterations
func IsLychrel(n int64, max_iterations int) bool {
	dig := Digits(n)
	for i := 0; i < max_iterations; i++ {

		dig2 := make([]byte, len(dig))
		copy(dig2, dig)
		DigitsReverse(dig2)

		dig = DigitsAdd(dig, dig2)

		if DigitsPalindromic(dig) {
			return false
		}
	}
	return true
}
