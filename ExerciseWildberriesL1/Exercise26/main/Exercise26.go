package main

import (
	"fmt"
	"strings"
)

func main() {
	var newString string
	_, err := fmt.Scan(&newString) // Добавляем новую строку для проверки
	checkError(err)

	var letters string // Создаем переменную для проверки уникальности, в которую будем сохранять буквы сразу в двух регистрах
	flag := true       // Создаем булеву переменную, если flag станет false - значит где-то был найден повтор

	for _, i := range newString { // Проверяем все буквы строки newString по одной
		/* Если переменная letters не пуста - сравниваем i с уже добавленными буквами и тем самым, узнаем нет ли повтора*/
		for _, o := range letters {
			if letters == "" {
				break
			} else if i == o {
				flag = false // Если повтор был - переменной flag присваивается false и заканчиваем цикл for для переменной letters
				break
			}
		}

		if flag == false { // Если flag == false, значит повтор был, заканчиваем цикл for для переменной newString
			break
		} else {
			/* Если повтора не было:
			1) Проверяем, какая буква нам дана(в верхнем регистре или нижнем)
			2) Исходя из пункта выше, мы прибавляем букву, и в верхнем регистре, и в нижнем, чтобы в будущих проверках
			для нас не имело значения в каком регистре находятся буквы*/
			if string(i) == strings.ToUpper(string(i)) {
				letters += string(i) + strings.ToLower(string(i))
			} else {
				letters += string(i) + strings.ToUpper(string(i))
			}
		}
	}
	// Выводим ответ в соответствии с условием:
	// Если повтор найден - false, если не найден - true
	if flag == false {
		fmt.Println("false")
	} else {
		fmt.Println("true")
	}
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
