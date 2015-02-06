package natural

import (
	"errors"
)

var romanValues map[rune]int

func init() {
	romanValues = make(map[rune]int, 16)
	romanValues['I'] = 1
	romanValues['V'] = 5
	romanValues['X'] = 10
	romanValues['L'] = 50
	romanValues['C'] = 100
	romanValues['D'] = 500
	romanValues['M'] = 1000
}

var ErrorInvalidRomanFormat = errors.New("Invalid roman numeral format")

// Parse a number in roman form
func ParseRoman(s string) (int, error) {

	num := 0
	lastVal := 99999
	for _, ch := range s {
		if val, found := romanValues[ch]; found {
			num += val
			if val > lastVal {
				num -= lastVal * 2
			}
			lastVal = val
		} else {
			return 0, ErrorInvalidRomanFormat
		}
	}

	return num, nil
}

// Format a number in the minimal roman format
func FormatRoman(num int) string {

	res := ""
	for num >= 1000 {
		res += "M"
		num -= 1000
	}

	if num >= 900 {
		res += "CM"
		num -= 900
	} else if num >= 500 {
		res += "D"
		num -= 500
	} else if num >= 400 {
		res += "CD"
		num -= 400
	}

	for num >= 100 {
		res += "C"
		num -= 100
	}

	if num >= 90 {
		res += "XC"
		num -= 90
	} else if num >= 50 {
		res += "L"
		num -= 50
	} else if num >= 40 {
		res += "XL"
		num -= 40
	}

	for num >= 10 {
		res += "X"
		num -= 10
	}

	if num >= 9 {
		res += "IX"
		num -= 9
	} else if num >= 5 {
		res += "V"
		num -= 5
	} else if num >= 4 {
		res += "IV"
		num -= 4
	}

	for num > 0 {
		res += "I"
		num -= 1
	}

	return res
}
