package natural

import "math/big"

func SqrtBig(n *big.Int) (x *big.Int) {
	switch n.Sign() {
	case -1:
		panic(-1)
	case 0:
		return big.NewInt(0)
	}

	var px, nx big.Int
	x = big.NewInt(0)
	x.SetBit(x, n.BitLen()/2+1, 1)
	for {
		nx.Rsh(nx.Add(x, nx.Div(n, x)), 1)
		if nx.Cmp(x) == 0 || nx.Cmp(&px) == 0 {
			break
		}
		px.Set(x)
		x.Set(&nx)
	}
	return
}
