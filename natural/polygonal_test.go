package natural

import "testing"

func testSerie(t *testing.T, gen func(int64) chan int64, check func(x int64) bool) {
	ch := gen(100000000)
	last := int64(1)
	for i := 0; i < 1000; i++ {
		n, ok := <-ch
		if !ok {
			t.Fatal("should not be closed")
		}
		if !check(n) {
			t.Fatal("should be:", n, i)
		}
		for j := last + 1; j < n; j++ {
			if check(j) {
				t.Fatal("should not be:", j, i)
			}
		}
		last = n
	}
}

func TestTriangular(t *testing.T) {
	testSerie(t, Triangular, IsTriangular)
}
func TestSquare(t *testing.T) {
	testSerie(t, Square, IsSquare)
}
func TestPentagonal(t *testing.T) {
	testSerie(t, Pentagonal, IsPentagonal)
}
func TestHexagonal(t *testing.T) {
	testSerie(t, Hexagonal, IsHexagonal)
}
func TestHeptagonal(t *testing.T) {
	testSerie(t, Heptagonal, IsHeptagonal)
}
func TestOctagonal(t *testing.T) {
	testSerie(t, Octagonal, IsOctagonal)
}
