package euler

import (
	"testing"

	"code.google.com/p/intmath/i64"
)

func TestP86(t *testing.T) {

	cnt := 0
	for i := int64(1); i <= 4000; i++ {
		for jk := int64(2); jk <= i+i; jk++ {
			psq := i*i + (jk)*(jk)
			p := i64.Sqrt(psq)
			if p*p == psq {
				// compute the number of combinations of j and k such as
				// j+k=jk and k<j<i
				if jk > i {
					cnt += int(2*i-jk)/2 + 1
				} else {
					cnt += int(jk / 2)
				}
			}
		}

		if i == 99 && cnt != 1975 {
			t.Error(i, cnt)
		} else if i == 100 && cnt != 2060 {
			t.Error(i, cnt)
		}

		if cnt >= 1000000 {
			if i != 1818 {
				t.Error(i, cnt)
			}
			return
		}
	}
}
