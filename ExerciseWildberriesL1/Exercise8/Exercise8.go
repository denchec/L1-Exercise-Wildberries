package main

import (
	"fmt"
	"time"
)

func main() {
	var timeValue time.Duration
	_, err := fmt.Scan(&timeValue)
	checkError(err)

	ch := make(chan int, 1)

	exitTime := time.After(time.Second * timeValue)

	var sequenceOfNumbers []int

	num := 1
	for {
		select {
		case <-exitTime:
			fmt.Println(len(sequenceOfNumbers))
			return
			/* Если время еще НЕ истекло, то:
			1) в канал передается число из переменной num
			2) число забирается из канала и добавляется в срез sequenceOfNumbers
			3) переменная num увеличивается на 1*/
		default:
			ch <- num
			sequenceOfNumbers = append(sequenceOfNumbers, <-ch)
			num += 1
		}
	}
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
