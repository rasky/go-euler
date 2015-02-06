package natural

import "testing"

func TestParseRoman(t *testing.T) {

	tests := []struct {
		s   string
		val int
	}{
		{"IL", 49},
		{"LIX", 59},
		{"XVII", 17},
		{"MDC", 1600},
		{"KX", -1},
	}

	for _, test := range tests {
		v, err := ParseRoman(test.s)
		if err != nil && test.val != -1 {
			t.Error("expected error", test.s)
		}
		if err == nil && test.val != v {
			t.Error(test.s, v)
		}
	}
}

func TestFormatRoman(t *testing.T) {

	tests := []struct {
		val int
		s   string
	}{
		{2959, "MMCMLIX"},
		{1744, "MDCCXLIV"},
		{1606, "MDCVI"},
		{49, "XLIX"},
		{888, "DCCCLXXXVIII"},
	}

	for _, test := range tests {
		s := FormatRoman(test.val)
		if s != test.s {
			t.Error(test.val, s)
		}
	}
}
