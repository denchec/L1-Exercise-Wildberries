package main

import (
	"fmt"
	"sync"
)

type SafeCounter struct {
	array []int
	mux   sync.Mutex
}

func main() {
	wg := sync.WaitGroup{}

	safeCount := SafeCounter{array: []int{2, 4, 6, 8, 10}}

	dataOfSquaring := make(chan int, len(safeCount.array))

	// Создали 5 горутин, которые подсчитывают квадрат чисел и записывают в канал
	for i := 0; i < len(safeCount.array); i++ {
		go func(i int) {
			wg.Add(1)
			defer wg.Done()
			safeCount.mux.Lock()
			dataOfSquaring <- safeCount.array[i] * safeCount.array[i]
			safeCount.mux.Unlock()
		}(i)
	}

	wg.Wait()

	var totalValue int
	for i := 0; i < len(safeCount.array); i++ {
		totalValue += <-dataOfSquaring
	}

	close(dataOfSquaring)

	fmt.Println(totalValue)
}
