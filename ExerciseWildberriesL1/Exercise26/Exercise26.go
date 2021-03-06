package main

import (
	"fmt"
	"strings"
)

func main() {
	var newString string
	_, err := fmt.Scan(&newString)
	checkError(err)

	var letters string
	flag := true

	for _, i := range newString {
		// Если переменная letters не пуста - сравниваем i с уже добавленными буквами и тем самым, узнаем нет ли повтора
		for _, o := range letters {
			if letters == "" {
				break
			} else if i == o {
				// Если повтор был, то переменной flag присваивается false и заканчиваем цикл для переменной letters
				flag = false
				break
			}
		}

		if flag == false { // Если flag == false, значит повтор был, заканчиваем цикл для переменной newString
			break
		} else {
			/* Если повтора не было:
			1) Проверяем, какая буква нам дана(в верхнем регистре или нижнем)
			2) Исходя из пункта выше, мы прибавляем букву, и в верхнем регистре, и в нижнем, чтобы в будущих проверках
			для нас не имело значения в каком регистре находятся буквы*/
			if string(i) == strings.ToUpper(string(i)) {
				letters += string(i) + strings.ToLower(string(i))
			} else {
				letters += string(i) + strings.ToUpper(string(i))
			}
		}
	}

	if flag == false {
		fmt.Println("false")
	} else {
		fmt.Println("true")
	}
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

//func main() {
//	strf := "aabcd"
//	res := true
//	m := make(map[string]struct{})
//
//	str := strings.ToLower(strf)
//
//	for _, c := range str {
//		m[string(c)] = struct{}{}
//	}
//
//	if len(m) < len(strf) {
//		res = false
//	}
//
//	fmt.Println(str+":", res)
//}
