package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Hello world!") // Выводим первое сообщение

	time.Sleep(2 * time.Second) // Ждем две секунды

	fmt.Println("It's me - Mario!") // Выводим второе сообщение
}
