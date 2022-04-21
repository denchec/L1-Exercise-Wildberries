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

	/* Если пользователь просит заменить 0 на 1:
	1) Сдвигаем единицу на необходимое количество символов, минус 1(это необходимо, чтобы изменить
	определенный символ), полученное число будет находиться в десятичной системе исчисления
	2) Прибавляем к первоначальному числу, получившееся в первом пункте */
	if dataMap[2] == 1 {
		endNum = dataMap[0] + dataMap[2]<<(dataMap[1]-1)
		/* Если пользователь просит заменить 1 на 0:
		Делаем, то же самое, что и при замене 0 на 1, только в конце вычитаем полученное число в результате сдвига
		единицы из первоначального */
	} else if dataMap[2] == 0 {
		endNum = dataMap[0] - 1<<(dataMap[1]-1)
	}

	fmt.Println(strconv.FormatInt(endNum, 2))
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
