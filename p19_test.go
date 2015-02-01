package euler

import "testing"

func p19MonthLen(year, month int) int {
	switch month {
	case 1, 3, 5, 7, 8, 10, 12:
		return 31
	case 4, 6, 9, 11:
		return 30
	case 2:
		leapyear := (year%4 == 0) && (!(year%100 == 0) || (year%400 == 0))
		if leapyear {
			return 29
		}
		return 28
	}
	panic("unreachable")
}

func TestP19(t *testing.T) {

	sundays := 0
	dow := 0
	for year := 1900; year < 2001; year++ {
		for month := 1; month <= 12; month++ {
			dow += p19MonthLen(year, month)
			dow = dow % 7
			if year == 1900 && month < 12 {
				continue
			}
			if dow == 6 {
				sundays++
			}
		}
	}

	if sundays != 171 {
		t.Error(sundays)
	}
}
