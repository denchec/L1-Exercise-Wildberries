package main

import "fmt"

func main() {
	newSlice := []int{1, 2, 9, 20, 31, 45, 63, 70, 100}

	fmt.Println(binarySearch(63, newSlice))
}

func binarySearch(num int, newSlice []int) bool {

	low := 0
	high := len(newSlice) - 1

	for low <= high {
		median := (low + high) / 2

		if newSlice[median] < num {
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
	if low == len(newSlice) || newSlice[low] != num {
		return false
	}

	return true
}
