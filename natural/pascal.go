package natural

import "math/big"

// Iterate over the specified rows (0-based) of the Pascal triangle
func PascalTriangle(row1, row2 int64) chan []*big.Int {

	if row1 < 0 || row1 >= row2 {
		panic("invalid argument")
	}

	ch := make(chan []*big.Int)
	go func() {
		defer close(ch)

		n := row1
		fact := FactorialArray(n + 1)
		row := make([]*big.Int, n+1)
		for r := int64(0); r <= (n+1)/2; r++ {
			row[r] = big.NewInt(0).Div(fact[n], big.NewInt(0).Mul(fact[r], fact[n-r]))
			row[n-r] = row[r]
		}

		ch <- row
		n += 1
		for n < row2 {
			newrow := make([]*big.Int, n+1)
			newrow[0] = big.NewInt(1)
			newrow[n] = newrow[0]
			if n >= 2 {
				newrow[1] = big.NewInt(n)
				newrow[n-1] = newrow[1]
			}
			for j := int64(2); j <= (n+1)/2; j++ {
				v := big.NewInt(0).Add(row[j-1], row[j])
				newrow[j] = v
				newrow[n-j] = newrow[j]
			}
			ch <- newrow
			row = newrow
			n += 1
		}
	}()
	return ch
}
