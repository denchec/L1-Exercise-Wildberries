package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Hello world!") // Выводим первое сообщение

	sleep(time.Second)

	fmt.Println("It's me - Mario!") // Выводим второе сообщение
}

func sleep(d time.Duration) {
	time.After(d)
}
