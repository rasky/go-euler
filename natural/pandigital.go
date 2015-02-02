package natural

// True if the specified digits are n-pandigital. Eg: n=9 means they are a
// permutation of 123456789.
// If n<0, n is equal to the number of specified digits (len(dig))
func PandigitalDigits(dig []byte, n int, zero bool) bool {
	if n < 0 {
		n = len(dig)
	}
	if zero && len(dig) != n+1 {
		return false
	} else if !zero && len(dig) != n {
		return false
	}

	var seen [10]bool
	if !zero {
		seen[0] = true // we don't want '0'
	}
	for _, d := range dig {
		if int(d) > n || seen[d] {
			return false
		}
		seen[d] = true
	}
	return true
}

// True if x is n-pandigital
func Pandigital(x int, n int, zero bool) bool {
	return PandigitalDigits(Digits(x), n, zero)
}

func PandigitalMulti(x []int, n int, zero bool) bool {
	dig := make([]byte, 0, 9)
	for _, x1 := range x {
		dig = append(dig, Digits(x1)...)
	}
	return PandigitalDigits(dig, n, zero)
}

func AllPandigitals(n int, zero bool) chan []byte {
	ch := make(chan []byte, 32)
	go func() {
		defer close(ch)
		dig := make([]byte, 0, 10)
		if zero {
			dig = append(dig, 0)
		}
		for i := 1; i <= n; i++ {
			dig = append(dig, byte(i))
		}

		ch <- dig
		for NextPermutation(dig) {
			dig2 := make([]byte, len(dig))
			copy(dig2, dig)
			ch <- dig2
		}
	}()
	return ch
}
