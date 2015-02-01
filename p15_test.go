package euler

import "testing"

type p15Size struct {
	w, h int
}
type p15NumPaths map[p15Size]int

func (p *p15NumPaths) setMemo(w, h int, num int) {
	if w > h {
		w, h = h, w
	}
	(*p)[p15Size{w, h}] = num
}

func (p *p15NumPaths) getMemo(w, h int) (int, bool) {
	if w == 0 || h == 0 {
		return 1, true
	}

	if w > h {
		w, h = h, w
	}
	num, found := (*p)[p15Size{w, h}]
	return num, found
}

func (p *p15NumPaths) Calc(w, h int) int {

	if num, found := p.getMemo(w, h); found {
		return num
	}

	num := 1
	for i := 0; i < w; i++ {
		num += p.Calc(w-i, h-1)
	}

	p.setMemo(w, h, num)
	return num
}

func TestP15(t *testing.T) {

	p := make(p15NumPaths)
	if p.Calc(1, 1) != 2 {
		t.Error(1, 1, p.Calc(1, 1))
	}
	if p.Calc(2, 1) != 3 {
		t.Error(2, 1, p.Calc(2, 1))
	}
	if p.Calc(2, 2) != 6 {
		t.Error(2, 2, p.Calc(2, 2))
	}
	if p.Calc(20, 20) != 137846528820 {
		t.Error(20, 20, p.Calc(20, 20))
	}
}
