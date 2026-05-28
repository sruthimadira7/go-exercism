package raindrops

import (
    "strconv"
    "strings"
)

var factors = []struct {
    divisor int
    sound string
}{
    {3, "Pling"},
    {5, "Plang"},
    {7, "Plong"},
}

func Convert(number int) string {
    var message strings.Builder

    for _, factor := range factors {
   		if number % factor.divisor == 0 {
            message.WriteString(factor.sound)
        }
    }

	if message.Len() == 0 {
        message.WriteString(strconv.Itoa(number))
    }

    return message.String()
}
