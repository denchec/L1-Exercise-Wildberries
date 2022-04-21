package main

import "fmt"

func main() {
	items := []int{1, 2, 9, 20, 31, 45, 63, 70, 100} // Создаем отсортированный массив

	fmt.Println(binarySearch(63, items))
}

func binarySearch(needle int, haystack []int) bool {

	low := 0                  // Объявляем переменную с индексом минимального значения массива
	high := len(haystack) - 1 // Объявляем переменную с индексом максимального значения массива

	for low <= high {
		median := (low + high) / 2 // Высчитываем середину массива.

		if haystack[median] < needle {
			/* Если искомое число больше середины массива, то мы присваиваем переменной low значение середины плюс 1,
			чтобы в следующий пробег цикла искать в половине, которая больше числа median. */
			low = median + 1
		} else {
			/* Если меньше, то мы присваиваем переменной high значение середины минус 1, чтобы в следующий пробег цикла
			искать в половине, которая меньше числа median. */
			high = median - 1
		}
	}

	// Если после исполнения цикла, не было найдено искомое значение - возвращаем false
	if low == len(haystack) || haystack[low] != needle {
		return false
	}

	return true
}
