package main

import (
	"fmt"
	"sync"
)

type SafeCounter struct {
	information map[string]string
	mux         sync.Mutex
}

func main() {
	mux := sync.RWMutex{}

	newMap := make(map[int]int)

	ch := make(chan int)

	for i := 1; i < 6; i++ {
		go func(i int) {
			mux.Lock()
			newMap[i] = i * i
			mux.Unlock()
			ch <- i
		}(i)
	}

	for i := 0; i < 5; i++ {
		// получаем из канала ключ по которому лежит значение
		x := <-ch
		fmt.Println("Ключ: ", x, "Значение: ", newMap[x])
	}
}
