package natural

// Return the Collatz sequence, starting from x
// The sequence is defined as:
//      n → n/2 (n is even)
//      n → 3n + 1 (n is odd)
// Although it has not been proved yet (Collatz Problem), it is thought
// that all starting numbers finish at 1.
func Collatz(x int32) chan int32 {
	ch := make(chan int32, 32)
	go func() {
		defer close(ch)
		for x != 1 {
			ch <- x
			if (x & 1) != 0 {
				x = 3*x + 1
			} else {
				x = x / 2
			}
		}
	}()
	return ch
}

var memoLenCollatz map[int64]int

func init() {
	memoLenCollatz = make(map[int64]int, 1000)
	memoLenCollatz[1] = 1
}

func LenCollatz(x int64) int {
	if x < 0 {
		panic("negative number")
	}

	l, found := memoLenCollatz[x]
	if found {
		return l
	}

	var next int64
	if (x & 1) != 0 {
		next = 3*x + 1
	} else {
		next = x / 2
	}

	l = LenCollatz(next) + 1
	memoLenCollatz[x] = l
	return l
}
