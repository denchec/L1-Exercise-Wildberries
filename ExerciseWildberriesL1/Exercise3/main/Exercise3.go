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
	wg := sync.WaitGroup{} // Переменная wg теперь отвечает за команды WaitGroup

	safeCount := SafeCounter{array: []int{2, 4, 6, 8, 10}} // положили последовательность чисел в array

	dataOfSquaring := make(chan int, len(safeCount.array)) // Открыли буферизированный канал, размер которого равен количеству чисел в последовательности

	// Создали 5 горутин, которые подсчитывают квадрат чисел и записывают в канал
	go func() {
		wg.Add(1)
		defer wg.Done()
		safeCount.mux.Lock()
		dataOfSquaring <- safeCount.array[0] * safeCount.array[0]
		safeCount.mux.Unlock()
	}()

	go func() {
		wg.Add(1)
		defer wg.Done()
		safeCount.mux.Lock()
		dataOfSquaring <- safeCount.array[1] * safeCount.array[1]
		safeCount.mux.Unlock()
	}()

	go func() {
		wg.Add(1)
		defer wg.Done()
		safeCount.mux.Lock()
		dataOfSquaring <- safeCount.array[2] * safeCount.array[2]
		safeCount.mux.Unlock()
	}()

	go func() {
		wg.Add(1)
		defer wg.Done()
		safeCount.mux.Lock()
		dataOfSquaring <- safeCount.array[3] * safeCount.array[3]
		safeCount.mux.Unlock()
	}()

	go func() {
		wg.Add(1)
		defer wg.Done()
		safeCount.mux.Lock()
		dataOfSquaring <- safeCount.array[4] * safeCount.array[4]
		safeCount.mux.Unlock()
	}()

	wg.Wait()

	var totalValue int
	// Через циклю for забираем все числа из канала и складываем
	for i := 0; i < len(safeCount.array); i++ {
		totalValue += <-dataOfSquaring
	}
	close(dataOfSquaring) // закрыли канал

	fmt.Println(totalValue)
}
