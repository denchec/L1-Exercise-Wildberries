package main

import (
	"fmt"
	"sync"
)

func main() {
	newSlice := []int{2, 4, 6, 8, 10}

	wg := sync.WaitGroup{}

	for i := 0; i < len(newSlice); i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			array := newSlice[i] * newSlice[i]
			fmt.Println(array)
		}(i)
	}

	wg.Wait()
}
