package main

import "fmt"

func main() {
	newMessageString := "snow dog sun"

	messageSlice := make(map[int]string)

	num := 0
	var word string

	for _, i := range newMessageString {

		if string(i) != " " {
			word += string(i)
		} else {
			messageSlice[num] = word
			word = ""
			num++
		}
	}
	messageSlice[num] = word // Записываем последнее слово в карту

	var endMessage string

	// Перебираем все значения в карте и прибавляем их к новой переменной, так же не забываем прибавить пробел
	for i := 0; i < len(messageSlice); i++ {
		endMessage = messageSlice[i] + " " + endMessage
	}

	fmt.Println(endMessage)
}
