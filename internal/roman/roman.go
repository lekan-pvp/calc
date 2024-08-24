package roman

import (
	"errors"
	"log"
	"strings"
)

var (
	ErrMoreThen10     = errors.New("the number more then 10")
	ErrLessThen1      = errors.New("the number less then 1")
	ErrEmptyString    = errors.New("the string is empty")
	ErrNotRomanNumber = errors.New("unknown roman number")
)

func FromInt(number int) (string, error) {
	var conversions = []struct {
		value   int
		numeral string
	}{
		{1000, "M"},
		{900, "CM"},
		{500, "D"},
		{400, "CD"},
		{100, "C"},
		{90, "XC"},
		{50, "L"},
		{40, "XL"},
		{10, "X"},
		{9, "IX"},
		{5, "V"},
		{4, "IV"},
		{1, "I"},
	}

	if number < 1 {
		return "", ErrLessThen1
	}

	var roman strings.Builder
	for _, conversion := range conversions {
		for number >= conversion.value {
			roman.WriteString(conversion.numeral)
			number -= conversion.value
		}
	}
	return roman.String(), nil
}

func FromRoman(s string) (int, error) {
	charMap := map[byte]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}

	length := len(s)
	if length == 0 {
		return 0, ErrEmptyString
	}

	sum := 0

	if length == 1 {
		sum, ok := charMap[s[0]]
		if !ok {
			return 0, ErrNotRomanNumber
		}
		if sum > 10 {
			return 0, ErrMoreThen10
		}
		return sum, nil
	}

	sum = charMap[s[length-1]]

	for i := length - 2; i >= 0; i-- {
		if v, ok := charMap[s[i]]; !ok {
			log.Println(v)
			return 0, ErrNotRomanNumber
		}
		if v, ok := charMap[s[i+1]]; !ok {
			log.Println(v)
			return 0, ErrNotRomanNumber
		}

		if charMap[s[i]] < charMap[s[i+1]] {
			sum -= charMap[s[i]]
		} else {
			sum += charMap[s[i]]
		}
	}

	if sum < 1 {
		return 0, ErrLessThen1
	}

	return sum, nil

}
