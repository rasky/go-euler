package euler

import (
	"euler/natural"
	"testing"
)

func TestP58(t *testing.T) {

	count := int64(0)

	n := int64(3)
	step := int64(2)
	for {
		if natural.IsPrime(n + step*0) {
			count += 1
		}
		if natural.IsPrime(n + step*1) {
			count += 1
		}
		if natural.IsPrime(n + step*2) {
			count += 1
		}
		if natural.IsPrime(n + step*3) {
			count += 1
		}

		num_diag := (step/2)*4 + 1
		if count*10/num_diag == 0 {
			break
		}

		n += step*4 + 2
		step += 2
	}

	if step+1 != 26241 {
		t.Error(step + 3)
	}
}
