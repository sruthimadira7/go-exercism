package raindrops

import (
    "strconv"
    "strings"
)

func addSound(message *strings.Builder, num int, divisor int, sound string) {
    if num % divisor == 0 {
        message.WriteString(sound)
    }
}


func Convert(number int) string {
    var message strings.Builder
	sound := map[int]string{
        3 : "Pling",
        5: "Plang",
        7: "Plong",
    }

    for i := 3; i<= 7; i += 2{
   		addSound(&message, number, i, sound[i])
    }

	if message.Len() == 0 {
        return strconv.Itoa(number)
    }

    return message.String()
}
