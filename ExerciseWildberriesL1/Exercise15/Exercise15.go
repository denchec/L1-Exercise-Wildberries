package main

import (
	"fmt"
	"strings"
)

/* Описание проблемы первоначального кода:

var justString string
func someFunc() {
  v := createHugeString(1 << 10)
  justString = v[:100]
}

func main() {
  someFunc()
}

1) Не хватает функции createHugeString, которая бы возвращала срез из 1024 элементов типа string.
2) В языке Golang существует garbage collector, который удаляет из памяти неиспользуемые значения для увеличения
быстродействия работы программы.
Из-за того, что мы присваиваем переменной justString только 100 значений, а не весь срез - 924 элемента болтаются
в памяти. При этом, garbage collector не может удалить неиспользуемые элементы т.к. срез justString ссылается на
основной срез.*/

// Решение вышеуказанных проблем:

var justString string

func someFunc() {
	// Создадим функцию, которая возвращает срез из 1024 строк
	v := createHugeString(1 << 10)
	/* Воспользуемся функцией Clone, которая копирует определенные элементы(в нашем случае - первые 100),
	и при этом не ссылается на общий срез, по этому garbage collector может удалить неиспользуемые элементы среза*/
	justString = strings.Clone(v[:100])
	fmt.Println(justString)
}

func main() {
	someFunc()
}

func createHugeString(a int) string {
	var b string
	for a > 0 {
		a--
		b += "qwerty"
	}
	return b
}
