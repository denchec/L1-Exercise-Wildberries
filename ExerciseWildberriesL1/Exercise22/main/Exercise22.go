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
	fmt.Println("A * B", C) // Умножение
	C = A + B
	fmt.Println("A + B", C) // Сложение
	C = A - B
	fmt.Println("A - B", C) // Разность
	C = A % B
	fmt.Println("A % B", C) // Вычисление остатка от деления
	C = A / B
	fmt.Println("A / B", C) // Целочисленное деление
	var D float64
	D = float64(A) / float64(B)
	fmt.Println("A / B", D) // Деление чисел с плавающей точкой
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
