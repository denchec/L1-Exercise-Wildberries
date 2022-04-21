package main

import "fmt"

func main() {
	newMessageString := "snow dog sun"

	messageSlice := make(map[int]string)

	num := 0
	var word string

	// Пробегаем по переменной newMessageString обращаясь к каждому символу отдельно
	for _, i := range newMessageString {

		if string(i) == " " { // Если встречаем символ пробела
			messageSlice[num] = word // Записываем переменную word в карту messageSlice с ключом num
			word = ""                // Очищаем переменную
			num += 1                 // Увеличиваем счетчик, чтобы при записи в карту - ключи отличались
		} else {
			word += string(i) // Складываем буквы с новой переменной
		}
	}
	messageSlice[num] = word // Записываем последнее слово в карту

	var endMessage string

	// Перебираем все значения в карте и прибавляем их к новой переменной, так же не забываем прибавить пробел
	for i := 0; i < len(messageSlice); i++ {
		endMessage = messageSlice[i] + " " + endMessage
	}

	fmt.Println(endMessage) // Выводим перевернутую строку
}
