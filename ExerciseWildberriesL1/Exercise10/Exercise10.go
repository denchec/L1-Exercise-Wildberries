package main

import "fmt"

func main() {
	arr := []float64{-25.4, -27.0, 13.0, -9.0, 15.5, 24.5, -21.0, 32.5}
	endMap := make(map[int][]float64)

	for _, temp := range arr {
		if temp < 0 {
			firstMapKey := (int(temp) - 10) / 10 * 10 // Объявляем переменную и записываем в нее ключ
			endMap[firstMapKey] = append(endMap[firstMapKey], temp)
		} else {
			secondMapKey := (int(temp) / 10) * 10                     // Объявляем переменную и записываем в нее ключ
			endMap[secondMapKey] = append(endMap[secondMapKey], temp) // Заполняем карту
		}
	}

	for k, v := range endMap {
		fmt.Println(k, v)
	}
}
