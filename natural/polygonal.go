package natural

import "code.google.com/p/intmath/i64"

type serieFunc func(i int64) int64

func serie(max int64, f serieFunc) chan int64 {
	ch := make(chan int64, 16)
	go func() {
		defer close(ch)
		i := int64(1)
		for {
			p := f(i)
			if p > max {
				return
			}
			ch <- p
			i += 1
		}
	}()
	return ch
}

// Return the serie of triangle numbers, generated by adding the natural numbers
// So the 7th triangle number would be 1 + 2 + 3 + 4 + 5 + 6 + 7 = 28.
func Triangular(max int64) chan int64 {
	return serie(max, func(n int64) int64 {
		return n * (n + 1) / 2
	})
}

func Square(max int64) chan int64 {
	return serie(max, func(n int64) int64 {
		return n * n
	})
}

func Pentagonal(max int64) chan int64 {
	return serie(max, func(n int64) int64 {
		return n * (3*n - 1) / 2
	})
}

func Hexagonal(max int64) chan int64 {
	return serie(max, func(n int64) int64 {
		return n * (2*n - 1)
	})
}

func Heptagonal(max int64) chan int64 {
	return serie(max, func(n int64) int64 {
		return n * (5*n - 3) / 2
	})
}

func Octagonal(max int64) chan int64 {
	return serie(max, func(n int64) int64 {
		return n * (3*n - 2)
	})
}

func IsTriangular(x int64) bool {
	sq := i64.Sqrt(8*x + 1)
	return sq*sq == 8*x+1 && ((sq-1)%2) == 0
}

func IsSquare(x int64) bool {
	sq := i64.Sqrt(x)
	return sq*sq == x
}

func IsPentagonal(x int64) bool {
	// Solve x = n(3n-1)/2 for n
	// 3n^2 -n -2x = 0
	// Has integer roots only if sqrt(b^2-4ac) is integer,
	// so sqrt(1 + 24*x)
	sq := i64.Sqrt(24*x + 1)
	return sq*sq == 24*x+1 && (sq%6) == 5
}

func IsHexagonal(x int64) bool {
	sq := i64.Sqrt(8*x + 1)
	return sq*sq == 8*x+1 && ((sq+1)%4) == 0
}

func IsHeptagonal(x int64) bool {
	sq := i64.Sqrt(40*x + 9)
	return sq*sq == 40*x+9 && ((sq+3)%10) == 0
}

func IsOctagonal(x int64) bool {
	sq := i64.Sqrt(3*x + 1)
	return sq*sq == 3*x+1 && ((sq+1)%3) == 0
}