package main

import (
	"fmt"
)

func main() {
	ar := []int{3, 7, 1, 2, 5, 4, -1, 0, 8}
	Quicksort(ar)
	fmt.Println(ar) // Вывели отсортированный массив
}

func Quicksort(ar []int) {
	if len(ar) <= 1 { // Если длина массива меньше или равна 1 - сортировка не нужна
		return
	}

	split := partition(ar) // Запустили функцию partition для сортировки массива

	Quicksort(ar[:split]) // Повторяем функцию Quicksort в левой части, от переменной split
	Quicksort(ar[split:]) // Повторяем функцию Quicksort в правой части, от переменной split
}

func partition(ar []int) int {
	center := ar[len(ar)/2] // Определяем середину массива

	left := 0            // Определяем переменную начала массива
	right := len(ar) - 1 // Определяем переменную конца массива

	/* Цикл For, в котором мы проверяем есть ли среди левых, от центра, значений те, что больше правых:
	1) Если переменные left и right встретились - то цикл останавливается, а переменная right выводится
	2) Если не происходит вышеуказанного пункта - меняем местами переменные ar[left] и ar[right] т.к. значение переменной ar[left] будет больше, чем ar[right] */
	for {
		for ; ar[left] < center; left++ {
		}

		for ; ar[right] > center; right-- {
		}

		if left >= right {
			return right
		}

		ar[left], ar[right] = ar[right], ar[left]
	}
}
