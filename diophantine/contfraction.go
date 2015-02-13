package diophantine

import (
	"fmt"
	"math/big"
	"reflect"
	"strconv"
	"strings"
)

type ContinuedFraction struct {
	FixedTerms  []int64
	PeriodTerms []int64
}

type triplet struct {
	a, b, c int64
}

// Computer floor(a/b), correctly computed also for negative results
// eg: -2/3 = -1
func sdiv_floor(n, d int64) int64 {
	sn := int64(1)
	if n < 0 {
		sn = -1
	}
	sd := int64(1)
	if d < 0 {
		sd = -1
	}
	if sn*sd > 0 {
		return n / d
	}
	return (n - (d*sd - 1)) / d
}

// Create a ContinuedFraction generated from the expansion of (sqrt(s)+m)/d
// If s<0, we compute the expansion of (-sqrt(abs(s))+m)/d
func NewContinuedFractionForSqrt(s, m, d int64) ContinuedFraction {

	cf := ContinuedFraction{}
	seen := make(map[triplet]int)

	sign := int64(1)
	if s < 0 {
		sign = -1
		s = -s
	}
	a0 := isqrt(s) * sign
	a := sdiv_floor(a0+m, d)
	cf.FixedTerms = append(cf.FixedTerms, a)
	seen[triplet{m, d, a}] = 0
	for idx := 1; ; idx++ {
		m = d*a - m
		d = (s - m*m) / d
		a = sdiv_floor(a0+m, d)
		if idx2, found := seen[triplet{m, d, a}]; found {
			cf.PeriodTerms = cf.FixedTerms[idx2:]
			cf.FixedTerms = cf.FixedTerms[:idx2]
			return cf
		}
		cf.FixedTerms = append(cf.FixedTerms, a)
		seen[triplet{m, d, a}] = idx
	}
}

func (cf ContinuedFraction) Equal(cf2 ContinuedFraction) bool {

	if !reflect.DeepEqual(cf.FixedTerms, cf2.FixedTerms) {
		return false
	}

	pt := make([]int64, len(cf2.PeriodTerms))
	copy(pt, cf2.PeriodTerms)
	for i := 0; i < len(pt); i++ {
		if reflect.DeepEqual(cf.PeriodTerms, pt) {
			return true
		}
		pt = append(pt[1:], pt[0])
	}
	return false
}

func (cf ContinuedFraction) String() string {

	fterms := make([]string, len(cf.FixedTerms))
	for idx, t := range cf.FixedTerms {
		fterms[idx] = strconv.Itoa(int(t))
	}

	pterms := make([]string, len(cf.PeriodTerms))
	for idx, t := range cf.PeriodTerms {
		pterms[idx] = strconv.Itoa(int(t))
	}

	var ap string
	if len(fterms) > 1 {
		ap = strings.Join(fterms[1:], ",") + ","
	}
	p := strings.Join(pterms, ",")
	return fmt.Sprintf("[%s; %s(%s)...]", fterms[0], ap, p)
}

func (cf ContinuedFraction) Evaluate(n int) *big.Rat {

	term := int64(0)
	next := func() bool {
		if n < 0 {
			return false
		}
		p := n - len(cf.FixedTerms)
		if p >= 0 {
			term = cf.PeriodTerms[p%len(cf.PeriodTerms)]
		} else {
			term = cf.FixedTerms[n]
		}
		n--
		return true
	}

	next()
	r := big.NewRat(term, 1)
	for next() {
		r.Inv(r)
		r.Add(r, big.NewRat(term, 1))
	}
	return r
}
