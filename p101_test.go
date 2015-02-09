package euler

import (
	"testing"

	"code.google.com/p/gomatrix/matrix"
	"code.google.com/p/intmath/intgr"
)

func TestP101(t *testing.T) {

	// Polygonal fit: find the best polynomial that fits the given points
	// Uses a Vandermonde matrix, solved through QR decomposition.
	polyfit := func(xin, yin []float64, n int) []float64 {
		m := len(yin)
		if len(yin) != len(xin) {
			panic("invalid args")
		}
		y := matrix.MakeDenseMatrix(yin, m, 1)
		x := matrix.Zeros(m, n)
		for i := 0; i < m; i++ {
			ip := float64(1)
			for j := 0; j < n; j++ {
				x.Set(i, j, ip)
				ip *= xin[i]
			}
		}

		q, r := x.QR()
		qty, err := q.Transpose().Times(y)
		if err != nil {
			panic(err)
		}
		c := make([]float64, n)
		for i := n - 1; i >= 0; i-- {
			c[i] = qty.Get(i, 0)
			for j := i + 1; j < n; j++ {
				c[i] -= c[j] * r.Get(i, j)
			}
			c[i] /= r.Get(i, i)
		}
		return c
	}

	polyfitint := func(y []int) []int {
		n := len(y)
		yin := make([]float64, n)
		xin := make([]float64, n)
		for i := 0; i < n; i++ {
			xin[i] = float64(i + 1)
			yin[i] = float64(y[i])
		}
		c := polyfit(xin, yin, n)
		ci := make([]int, n)
		for i := 0; i < n; i++ {
			if c[i] < 0 {
				ci[i] = int(c[i] - .5)
			} else {
				ci[i] = int(c[i] + .5)
			}
		}
		return ci
	}

	sumfits := func(y []int, n int) int {
		sum := 0
		for i := 1; i <= n; i++ {
			c := polyfitint(y[:i])

			// Compute the FIT (first incorrect term)
			v := c[0]
			ip := i + 1
			for j := 1; j < i; j++ {
				v += ip * c[j]
				ip *= (i + 1)
			}

			sum += v
		}
		return sum
	}

	y := []int{1, 8, 27, 64, 125, 216}
	if sumfits(y, 3) != 74 {
		t.Error(sumfits(y, 3))
	}

	y = make([]int, 10)
	for i := 1; i < len(y)+1; i++ {
		y[i-1] = 1 - i + intgr.Pow(i, 2) - intgr.Pow(i, 3) + intgr.Pow(i, 4) - intgr.Pow(i, 5) + intgr.Pow(i, 6) - intgr.Pow(i, 7) + intgr.Pow(i, 8) - intgr.Pow(i, 9) + intgr.Pow(i, 10)
	}
	if sumfits(y, 10) != 37076114526 {
		t.Error(sumfits(y, 10))
	}

}
