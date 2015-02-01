package natural

import "math/big"

func Fibo() chan int {
	ch := make(chan int)

	go func() {
		i := 1
		j := 2

		for {
			ch <- i
			k := i + j
			i = j
			j = k
		}
	}()

	return ch
}

func FiboBig() chan *big.Int {
	ch := make(chan *big.Int)

	go func() {
		i := big.NewInt(1)
		j := big.NewInt(2)

		for {
			ch <- i
			k := big.NewInt(0).Add(i, j)
			i = j
			j = k
		}
	}()

	return ch
}
