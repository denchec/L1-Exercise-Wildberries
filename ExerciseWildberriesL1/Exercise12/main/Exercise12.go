package main

import "fmt"

func main() {
	var newScan string // Создаем переменную, в которую будет записываться строка с клавиатуры

	var array []string // Создаем переменную для записи множества

	for i := 0; i < 5; i++ { // Создаем цикл for в 5 шагов
		fmt.Println("Введите строку:")
		_, err := fmt.Scan(&newScan) // Сканируем введенную строку
		checkError(err)
		array = append(array, newScan) // Записываем новую строку в переменную множества
	}

	fmt.Println(array) // Выводим множество
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
