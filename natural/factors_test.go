package natural

import (
	"reflect"
	"testing"
)

func allFactors(x int64) []int32 {
	fact := make([]int32, 0, 8)
	for f := range Factors(x) {
		fact = append(fact, f)
	}
	return fact
}

func TestPrimes(t *testing.T) {

	primes := make([]int32, 0, 32)
	for p := range Primes(1000) {
		primes = append(primes, p)
	}
	exp := []int32{
		2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41, 43, 47, 53, 59, 61, 67, 71,
		73, 79, 83, 89, 97, 101, 103, 107, 109, 113, 127, 131, 137, 139, 149, 151, 157, 163, 167, 173,
		179, 181, 191, 193, 197, 199, 211, 223, 227, 229, 233, 239, 241, 251, 257, 263, 269, 271, 277, 281,
		283, 293, 307, 311, 313, 317, 331, 337, 347, 349, 353, 359, 367, 373, 379, 383, 389, 397, 401, 409,
		419, 421, 431, 433, 439, 443, 449, 457, 461, 463, 467, 479, 487, 491, 499, 503, 509, 521, 523, 541,
		547, 557, 563, 569, 571, 577, 587, 593, 599, 601, 607, 613, 617, 619, 631, 641, 643, 647, 653, 659,
		661, 673, 677, 683, 691, 701, 709, 719, 727, 733, 739, 743, 751, 757, 761, 769, 773, 787, 797, 809,
		811, 821, 823, 827, 829, 839, 853, 857, 859, 863, 877, 881, 883, 887, 907, 911, 919, 929, 937, 941,
		947, 953, 967, 971, 977, 983, 991, 997}
	if !reflect.DeepEqual(primes, exp) {
		t.Error("invalid primes", primes)
	}
}

func TestFactors(t *testing.T) {

	fact := allFactors(13195)
	exp := []int32{5, 7, 13, 29}
	if !reflect.DeepEqual(fact, exp) {
		t.Error("invalid factors", fact)
	}

	fact = allFactors(14)
	exp = []int32{2, 7}
	if !reflect.DeepEqual(fact, exp) {
		t.Error("invalid factors", fact)
	}

	fact = allFactors(48)
	exp = []int32{2, 2, 2, 2, 3}
	if !reflect.DeepEqual(fact, exp) {
		t.Error("invalid factors", fact)
	}
}

func TestNumDivisors(t *testing.T) {

	if NumDivisors(28) != 6 {
		t.Error(NumDivisors(28))
	}
}
