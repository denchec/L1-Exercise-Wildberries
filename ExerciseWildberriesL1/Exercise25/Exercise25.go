package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Hello world!")

	sleep(time.Second)

	fmt.Println("It's me - Mario!")
}

func sleep(d time.Duration) {
	time.After(d)
}
