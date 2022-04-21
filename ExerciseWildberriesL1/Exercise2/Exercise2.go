package main

import (
	"fmt"
	"os"
	"sync"
)

type SafeCounter struct {
	arrayOfNumbers []int
	mux            sync.Mutex
}

func main() {
	safeCount := SafeCounter{arrayOfNumbers: []int{2, 4, 6, 8, 10}}

	wg := sync.WaitGroup{}

	wg.Add(1)
	go func() {
		defer wg.Done()
		safeCount.mux.Lock()
		array := safeCount.arrayOfNumbers[0] * safeCount.arrayOfNumbers[0]
		_, err := fmt.Fprintln(os.Stdout, array)
		checkError(err)
		safeCount.mux.Unlock()
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		safeCount.mux.Lock()
		array := safeCount.arrayOfNumbers[1] * safeCount.arrayOfNumbers[1]
		_, err := fmt.Fprintln(os.Stdout, array)
		checkError(err)
		safeCount.mux.Unlock()
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		safeCount.mux.Lock()
		array := safeCount.arrayOfNumbers[2] * safeCount.arrayOfNumbers[2]
		_, err := fmt.Fprintln(os.Stdout, array)
		checkError(err)
		safeCount.mux.Unlock()
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		safeCount.mux.Lock()
		array := safeCount.arrayOfNumbers[3] * safeCount.arrayOfNumbers[3]
		_, err := fmt.Fprintln(os.Stdout, array)
		checkError(err)
		safeCount.mux.Unlock()
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		safeCount.mux.Lock()
		array := safeCount.arrayOfNumbers[4] * safeCount.arrayOfNumbers[4]
		_, err := fmt.Fprintln(os.Stdout, array)
		checkError(err)
		safeCount.mux.Unlock()
	}()

	wg.Wait()
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
