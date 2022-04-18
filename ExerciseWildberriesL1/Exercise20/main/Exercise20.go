package main

import "fmt"

func main() {
	newMessageString := "snow dog sun"

	messageSlice := make(map[int]string)

	num := 1
	var word string

	for _, i := range newMessageString {
		if string(i) == " " {
			messageSlice[num] = word
			word = ""
			num += 1
		} else {
			word += string(i)
		}
	}
	messageSlice[num] = word

	var endMessage string

	for i := 1; i < len(messageSlice)+1; i++ {
		endMessage = messageSlice[i] + " " + endMessage
	}

	fmt.Println(endMessage)
}
