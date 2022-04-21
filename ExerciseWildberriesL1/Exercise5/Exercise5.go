package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var numSec time.Duration
	_, err := fmt.Scan(&numSec)
	checkError(err)

	wg := sync.WaitGroup{}

	ch := make(chan int, 1)

	exitTime := time.After(time.Second * numSec)

	sequenceOfNumbers := []int{}

	wg.Add(1)
	go func() {
		num := 1
		for {
			select {
			case <-exitTime:
				wg.Done()
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
	}()

	wg.Wait()
	fmt.Println(len(sequenceOfNumbers)) // Выводим количество новых записей, созданных за отведенный промежуток времени
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
