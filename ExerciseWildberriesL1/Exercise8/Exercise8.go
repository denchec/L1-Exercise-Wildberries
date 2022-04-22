package main

import (
	"fmt"
	"strconv"
)

func main() {
	dataMap := make(map[int]int64)

	// Не смог додуматься, как сканировать сразу в карту, пришлось добавить эту переменную
	var helpScan int64

	// Заполняем карту введенными данными с клавиатуры, присваивая i в качестве ключа
	for i := 0; i < 3; i++ {
		_, err := fmt.Scan(dataMap[i])
		checkError(err)
		dataMap[i] = helpScan
	}

	fmt.Println(strconv.FormatInt(dataMap[0], 2))

	var endNum int64

	if dataMap[2] == 1 {
		endNum = dataMap[0] | dataMap[2]<<(dataMap[1]-1)
	} else if dataMap[2] == 0 {
		endNum = dataMap[0] & (9223372036854775807 - 1<<(dataMap[1]-1))
	}

	fmt.Println(strconv.FormatInt(endNum, 2))
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
