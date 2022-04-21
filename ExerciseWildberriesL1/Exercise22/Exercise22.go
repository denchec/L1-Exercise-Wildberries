package main

import "fmt"

func main() {
	var A int
	var B int

	_, aErr := fmt.Scan(&A)
	checkError(aErr)
	_, bErr := fmt.Scan(&B)
	checkError(bErr)

	C := A * B
	fmt.Println("A * B", C) // Выводим значение умножения
	C = A + B
	fmt.Println("A + B", C) // Выводим значение сложения
	C = A - B
	fmt.Println("A - B", C) // Выводим значение разности
	C = A % B
	fmt.Println("A % B", C) // Выводим значение вычисления остатка от деления
	C = A / B
	fmt.Println("A / B", C) // Выводим значение целочисленного деления
	var D float64
	D = float64(A) / float64(B)
	fmt.Println("A / B", D) // Выводим значение деления чисел с плавающей точкой
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
