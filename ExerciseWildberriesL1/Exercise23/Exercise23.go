package main

import "fmt"

func main() {
	firstSlice := []int{1, 2, 3, 4, 5}

	var i int
	_, err := fmt.Scan(&i)
	checkError(err)

	firstWay(firstSlice, i)
	secondWay(firstSlice, i)
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

func firstWay(firstSlice []int, i int) {
	var newSliceA []int

	/* Проходим по всему срезу firstSlice и сверяем с введенным значением i,
	если число не является введенным значением - записываем его в новый срез*/
	for num := 0; num < len(firstSlice); num++ {
		if firstSlice[num] != firstSlice[i] {
			newSliceA = append(newSliceA, firstSlice[num])
		}
	}

	fmt.Println(newSliceA)
}

func secondWay(firstSlice []int, i int) {
	var newSliceB []int

	/* Записываем в новый срез все до введенного числа и все после*/
	newSliceB = append(firstSlice[:i], firstSlice[i+1:]...)

	fmt.Println(newSliceB)
}
