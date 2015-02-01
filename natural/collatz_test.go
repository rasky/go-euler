package natural

import "testing"

func TestLenCollatz(t *testing.T) {
	if LenCollatz(13) != 10 {
		t.Error(LenCollatz(13))
	}

	if LenCollatz(113383) != 248 {
		t.Error(LenCollatz(113383))
	}
}
