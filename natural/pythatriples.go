package natural

func PythagoreanTriplesWithSum(sum int) chan [3]int {
	ch := make(chan [3]int, 16)
	go func() {
		defer close(ch)
		for a := 2; a < sum; a++ {
			for b := a; b <= (sum-a)/2; b++ {
				c := sum - a - b
				if a*a+b*b == c*c {
					ch <- [3]int{a, b, c}
				}
			}
		}
	}()
	return ch
}
