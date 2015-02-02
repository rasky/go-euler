package euler

import (
	"euler/natural"
	"testing"
)

func TestP43(t *testing.T) {

	sum := int64(0)

	for dig := range natural.AllPandigitals(9, true) {
		if dig[3]%2 != 0 {
			continue
		}
		if (dig[2]+dig[3]+dig[4])%3 != 0 {
			continue
		}
		if dig[5] != 5 && dig[5] != 0 {
			continue
		}
		if (int(dig[4])*100+int(dig[5])*10+int(dig[6]))%7 != 0 {
			continue
		}
		if (int(dig[5])*100+int(dig[6])*10+int(dig[7]))%11 != 0 {
			continue
		}
		if (int(dig[6])*100+int(dig[7])*10+int(dig[8]))%13 != 0 {
			continue
		}
		if (int(dig[7])*100+int(dig[8])*10+int(dig[9]))%17 != 0 {
			continue
		}

		x := natural.FromDigitsReverse(dig)
		sum += int64(x)
	}

	if sum != 16695334890 {
		t.Error(sum)
	}
}
