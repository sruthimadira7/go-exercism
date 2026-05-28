package isbnverifier

import (
	"strings"
)

func IsValidISBN(isbn string) bool {
	isbn = strings.ReplaceAll(isbn, "-", "")

	if len(isbn) != 10 {
		return false
	}

	sum := 0
	for i := 0; i < 10; i++ {
		char := isbn[i]
		var digit int

		if char >= '0' && char <= '9' {
			digit = int(char - '0')
		} else if char == 'X' && i == 9 {
			digit = 10
		} else {
			return false
		}
		sum += digit * (10 - i)
	}
	return sum%11 == 0
}
