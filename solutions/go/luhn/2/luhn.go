package luhn

import (
    "strings"
    "unicode"
)

func Valid(id string) bool {
	id = strings.ReplaceAll(id, " ", "")

	if len(id) <= 1 {
		return false
	}

	total_sum := 0
	double := false
	for i := len(id) - 1; i >= 0; i-- {
		r := rune(id[i])
		if !unicode.IsDigit(r) {
			return false
		}
		digit := int(r - '0')
		if double {
			digit *= 2
			if digit > 9 {
				digit -= 9
			}
		}
		total_sum += digit
		double = !double
	}

	return total_sum%10 == 0
}