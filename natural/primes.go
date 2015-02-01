package natural

import "errors"

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

var sieve *BitArray
var lastSieveFactor int32

const maxPrime = 100000000

var allPrimes []int32

func init() {
	sieve = NewBitArray(maxPrime)
	allPrimes = make([]int32, 1, 1000000)
	allPrimes[0] = 2
	lastSieveFactor = 3
}

func Primes(max int32) chan int32 {
	if max > maxPrime {
		panic(errors.New("requested too big prime for sieve"))
	}

	ch := make(chan int32, 16)
	go func() {
		defer close(ch)

		for idx := 0; idx < len(allPrimes); idx++ {
			p := allPrimes[idx]
			if p > max {
				return
			}
			ch <- p
		}

		for y := lastSieveFactor; y <= max; y += 2 {
			if sieve.Get(y) {
				continue
			}

			for z := y + y; z <= max; z += y {
				sieve.Set(z)
			}
			lastSieveFactor = y

			allPrimes = append(allPrimes, y)
			ch <- y
		}
	}()
	return ch
}
