package natural

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
