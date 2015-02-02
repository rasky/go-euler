package natural

import "testing"

func TestLychrel(t *testing.T) {
	if !IsLychrel(196, 1000) {
		t.Error("196")
	}
	if IsLychrel(47, 1000) {
		t.Error("47")
	}
}
