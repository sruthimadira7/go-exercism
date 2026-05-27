package raindrops

import (
    "strconv"
    "strings")

func addSound(message *strings.Builder, num int, divisor int, sound string) {
    if num % divisor == 0 {
        message.WriteString(sound)
    }
}


func Convert(number int) string {
    //  struct type
    var message strings.Builder

	//  structs are copied by values
    //  need to pass areference
   addSound(&message, number, 3, "Pling")
   addSound(&message, number, 5, "Plang")
   addSound(&message, number, 7, "Plong")

	if message.Len() == 0 {
        return strconv.Itoa(number)
    }

    return message.String()
}
