package main

import (
	"fmt"
	"reflect"
)

func main() {
	var valueInterface interface{} = 1          // Объявляем переменную типа интерфейс со значением integer
	fmt.Println(reflect.TypeOf(valueInterface)) // Выводим тип переменной интерфейса - int
	valueInterface = "Hi"
	fmt.Println(reflect.TypeOf(valueInterface)) // Выводим тип переменной интерфейса - string
	valueInterface = true
	fmt.Println(reflect.TypeOf(valueInterface)) // Выводим тип переменной интерфейса - bool
	valueInterface = make(chan int)
	fmt.Println(reflect.TypeOf(valueInterface)) // Выводим тип переменной интерфейса - channel
}
