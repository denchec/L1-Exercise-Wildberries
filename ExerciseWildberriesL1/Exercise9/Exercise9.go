package main

import "fmt"

func main() {
	m := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	firstChannel := make(chan int, 1)
	secondChannel := make(chan int, 1)
	boolChannel := make(chan bool, 1)

	go func() {
		for _, i := range m {
			firstChannel <- i
		}
		close(firstChannel)
	}()

	go func() {
		for num := range firstChannel {
			square := num * num
			secondChannel <- square
		}
		close(secondChannel)
	}()

	go func() {
		for num := range secondChannel {
			fmt.Println(num)
		}
		boolChannel <- true
	}()
	<-boolChannel // Горутина main ждет значения bool
}
