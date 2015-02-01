package natural

import (
	"errors"

	"code.google.com/p/intmath/i64"
)

type BitArray struct {
	bits []uint8
}

func NewBitArray(sz int) *BitArray {
	return &BitArray{
		bits: make([]uint8, (sz+7)/8),
	}
}

func (b *BitArray) Set(idx int32) {
	b.bits[idx/8] |= 1 << (uint(idx) & 7)
}

func (b *BitArray) Get(idx int32) bool {
	return (b.bits[idx/8] & (1 << (uint(idx) & 7))) != 0
}

func (b *BitArray) Shrink(idx int) {
	if idx > len(b.bits) {
		panic(errors.New("not shrinking"))
	}
	b.bits = b.bits[0 : idx/8]
}

func Primes(max int32) chan int32 {
	ch := make(chan int32)
	go func() {
		defer close(ch)

		sieve := NewBitArray(int(max + 1))
		ch <- 2
		for y := int32(3); y <= max; y += 2 {
			if sieve.Get(y) {
				continue
			}
			ch <- y
			for z := y; z <= max; z += y {
				sieve.Set(z)
			}
		}
	}()
	return ch
}

func Factors(x int64) chan int32 {

	ch := make(chan int32)

	go func() {
		defer close(ch)

		limit := i64.Sqrt(x)

		if x&1 == 0 {
			x = x >> 1
			limit = i64.Sqrt(x)
			ch <- 2
		}

		for y := range Primes(int32(limit)) {
			if (x % int64(y)) == 0 {
				x = x / int64(y)
				limit = i64.Sqrt(x)
				ch <- y
			}
		}
	}()

	return ch
}
