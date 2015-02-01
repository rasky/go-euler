package euler

import "testing"

func TestP17(t *testing.T) {

	base := []int{
		len("zero"),
		len("one"),
		len("two"),
		len("three"),
		len("four"),
		len("five"),
		len("six"),
		len("seven"),
		len("eight"),
		len("nine"),
		len("ten"),
		len("eleven"),
		len("twelve"),
		len("thirteen"),
		len("fourteen"),
		len("fifteen"),
		len("sixteen"),
		len("seventeen"),
		len("eighteen"),
		len("nineteen"),
	}

	decades := []int{
		len("twenty"),
		len("thirty"),
		len("forty"),
		len("fifty"),
		len("sixty"),
		len("seventy"),
		len("eighty"),
		len("ninety"),
	}

	var englen func(int) int
	englen = func(x int) int {
		if x < 20 {
			return base[x]
		}
		if x < 100 {
			dec := x / 10
			ret := decades[dec-2]
			if x%10 != 0 {
				ret += base[x%10]
			}
			return ret
		}
		if x < 1000 {
			cent := x / 100
			ret := base[cent] + len("hundred")
			if x%100 != 0 {
				ret += len("and") + englen(x%100)
			}
			return ret
		}
		if x == 1000 {
			return len("one") + len("thousand")
		}
		panic("too big")
	}

	if englen(42) != 8 {
		t.Error(englen(42))
	}
	if englen(342) != 23 {
		t.Error(englen(342))
	}
	if englen(115) != 20 {
		t.Error(englen(115))
	}

	total := 0
	for i := 1; i <= 1000; i++ {
		total += englen(i)
	}
	if total != 21124 {
		t.Error(total)
	}
}
