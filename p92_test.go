package euler

import "testing"

func TestP92(t *testing.T) {

	cnt := 0
	for n := int64(1); n < 10000000; n++ {
		m := n
		for m != 1 && m != 89 {
			next := int64(0)
			for m > 0 {
				next += (m % 10) * (m % 10)
				m /= 10
			}
			m = next
		}
		if m == 89 {
			cnt++
		}
	}

	if cnt != 8581146 {
		t.Error(cnt)
	}
}
