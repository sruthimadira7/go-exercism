package raindrops

import (
    "strconv"
    "strings"
)

type Factor struct {
    divisor int
    sound string
}

func Convert(number int) string {
    var message strings.Builder
	factors := []Factor{
        {3, "Pling"},
        {5, "Plang"},
        {7, "Plong"},
    }

    for _, factor := range factors{
   		if number % factor.divisor == 0 {
            message.WriteString(factor.sound)
        }
    }

	if message.Len() == 0 {
        message.WriteString(strconv.Itoa(number))
    }

    return message.String()
}
