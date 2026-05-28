package luhn

import (
    "strings"
    "unicode"
)

func Valid(id string) bool {
	id = strings.Join(strings.Split(id, " "), "")

	if len(id) <= 1 {
		return false
	}

	total_sum := 0
	cnt := 0
	for i := len(id) - 1; i >= 0; i-- {
		r := rune(id[i])
		if !unicode.IsDigit(r) {
			return false
		}
		cnt += 1
		digit := int(id[i] - '0')
		doubleDigit := 0
		if cnt%2 == 0 {
			doubleDigit = digit * 2
			if doubleDigit > 9 {
				doubleDigit -= 9
			}
		} else {
			doubleDigit = digit
		}
		total_sum += doubleDigit

	}

	return total_sum%10 == 0
}
