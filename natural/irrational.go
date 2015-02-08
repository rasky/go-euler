package natural

import "math/big"

// Calculate the square root of n, up to prec digits. The channel returns
// the digits one by one, with the decimal point represented by 0xFF
func Sqrt(n int64, prec int) <-chan byte {

	ch := make(chan byte, 8)

	go func() {
		defer close(ch)

		dig := Digits(n)
		if len(dig)%2 != 0 {
			dig = append(dig, 0)
		}
		pidx := len(dig) / 2

		zero := big.NewInt(0)
		twenty := big.NewInt(20)
		root := new(big.Int)
		carry := new(big.Int)
		for i := 0; i < prec; i++ {
			carry.Mul(carry, big.NewInt(100))
			if len(dig) > 0 {
				n := len(dig)
				carry.Add(carry, big.NewInt(int64(dig[n-1])*10+int64(dig[n-2])))
				dig = dig[:n-2]
			}
			x := int64(9)
			root20 := new(big.Int).Mul(root, twenty)
			tmp := new(big.Int)
			bx := big.NewInt(x)
			for {
				tmp.Add(root20, bx).Mul(tmp, bx)
				if tmp.Cmp(carry) <= 0 {
					break
				}
				x -= 1
				bx = big.NewInt(x)
			}
			carry.Sub(carry, tmp)
			root.Mul(root, big.NewInt(10)).Add(root, bx)
			ch <- byte(x)
			if len(dig) == 0 && carry.Cmp(zero) == 0 {
				break
			}
			if pidx > 0 {
				pidx--
				if pidx == 0 {
					ch <- 0xFF
				}
			}
		}
	}()
	return ch
}
