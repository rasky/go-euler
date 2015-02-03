package euler

import (
	"euler/natural"

	"testing"
)

func TestP61(t *testing.T) {

	var buf [10000]byte
	for i := int64(1000); i < 10000; i++ {
		count := 0
		if natural.IsTriangular(i) {
			buf[i] |= 1
			count += 1
		}
		if natural.IsSquare(i) {
			buf[i] |= 2
			count += 1
		}
		if natural.IsPentagonal(i) {
			buf[i] |= 4
			count += 1
		}
		if natural.IsHexagonal(i) {
			buf[i] |= 8
			count += 1
		}
		if natural.IsHeptagonal(i) {
			buf[i] |= 16
			count += 1
		}
		if natural.IsOctagonal(i) {
			buf[i] |= 32
			count += 1
		}
	}

	for ax := 10; ax < 100; ax++ {
		for bx := 10; bx < 100; bx++ {
			if !(bx != ax) {
				continue
			}
			bc := buf[ax*100+bx]
			if bc == 0 {
				continue
			}
			for cx := 10; cx < 100; cx++ {
				if !(cx != bx && cx != ax) {
					continue
				}
				cc := bc | buf[bx*100+cx]
				if cc == bc {
					continue
				}
				for dx := 10; dx < 100; dx++ {
					if !(dx != cx && dx != bx && dx != ax) {
						continue
					}
					dc := cc | buf[cx*100+dx]
					if dc == cc {
						continue
					}
					for ex := 10; ex < 100; ex++ {
						if !(ex != dx && ex != cx && ex != bx && ex != ax) {
							continue
						}
						ec := dc | buf[dx*100+ex]
						if ec == dc {
							continue
						}

						for fx := 10; fx < 100; fx++ {
							if !(fx != ex && fx != dx && fx != cx && fx != bx && fx != ax) {
								continue
							}
							fc := ec | buf[ex*100+fx]
							if fc == ec {
								continue
							}

							gc := fc | buf[fx*100+ax]
							if gc == fc {
								continue
							}

							solution := []int64{
								int64(ax)*100 + int64(bx),
								int64(bx)*100 + int64(cx),
								int64(cx)*100 + int64(dx),
								int64(dx)*100 + int64(ex),
								int64(ex)*100 + int64(fx),
								int64(fx)*100 + int64(ax),
							}

							for _, s := range solution {
								t.Log(s, buf[s])
							}

							if natural.Sum(solution) == 28684 {
								return
							}
						}
					}
				}
			}
		}
	}

	t.Fail()
}
