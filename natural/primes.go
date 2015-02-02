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

func (b *BitArray) Set(idx int64) {
	b.bits[idx/8] |= 1 << (uint(idx) & 7)
}

func (b *BitArray) Get(idx int64) bool {
	return (b.bits[idx/8] & (1 << (uint(idx) & 7))) != 0
}

func (b *BitArray) Shrink(idx int) {
	if idx > len(b.bits) {
		panic(errors.New("not shrinking"))
	}
	b.bits = b.bits[0 : idx/8]
}

var sieve *BitArray
var lastSieveFactor int64

const maxPrime = 2 * 1000 * 1000

var allPrimes []int64

func init() {
	sieve = NewBitArray(maxPrime)
	allPrimes = make([]int64, 1, 1000000)
	allPrimes[0] = 2
	lastSieveFactor = int64(1)
}

func Primes(max int64) chan int64 {
	if max > maxPrime {
		panic(errors.New("requested too big prime for sieve"))
	}

	ch := make(chan int64, 16)
	go func() {
		defer close(ch)

		for idx := 0; idx < len(allPrimes); idx++ {
			p := allPrimes[idx]
			if p > max {
				return
			}
			ch <- p
		}

		for y := lastSieveFactor + 2; y <= max; y += 2 {
			if sieve.Get(y) {
				continue
			}

			for z := y + y; z < maxPrime; z += y {
				sieve.Set(z)
			}
			lastSieveFactor = y

			allPrimes = append(allPrimes, y)
			ch <- y
		}
	}()
	return ch
}

func IsPrime(p int64) bool {
	if p <= 3 {
		return p >= 2
	}
	if p%2 == 0 || p%3 == 0 {
		return false
	}
	for i := int64(5); i*i <= p; i += 6 {
		if p%i == 0 || p%(i+2) == 0 {
			return false
		}
	}
	return true
}
