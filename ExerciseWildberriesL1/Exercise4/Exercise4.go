package main

import (
	"fmt"
	"os"
	"os/signal"
	"strings"
	"sync"
	"syscall"
)

func main() {
	var amountOfWorkers int
	_, err := fmt.Scan(&amountOfWorkers)
	checkError(err)

	var wg sync.WaitGroup

	channel := make(chan string, amountOfWorkers)

	// Открываем канал в который будет подаваться сигнал о завершении программы
	sigs := make(chan os.Signal, 1)
	// говорим каналу sigs, чтобы он ждал только сигнал завершения программы
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	for i := 1; i < amountOfWorkers+1; i++ {
		newWord := strings.Repeat("x", i) + "HELLO" + strings.Repeat("x", i)
		channel <- newWord
		wg.Add(1)
		go func() {
			defer wg.Done()
			for {
				select {
				/* Если сигнал о завершении программы будет получен, то:
				1) мы запишем данные сигнала в переменную signalSigterm
				2) вернем сигнал из переменной signalSigterm обратно в канал,
					чтобы другие горутины могли его использовать
				3) return завершит горутину*/
				case signalSigterm := <-sigs:
					sigs <- signalSigterm
					return
				default:
					sayHello := <-channel
					_, err := fmt.Fprintln(os.Stdout, sayHello)
					checkError(err)
					channel <- sayHello
				}
			}
		}()
	}

	wg.Wait()
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
