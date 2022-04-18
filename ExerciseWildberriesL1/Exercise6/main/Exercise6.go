package main

import (
	"context"
	"errors"
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

func main() {
	/* Обобщение в функциях: Все следующие функции показывают разные варианты закрытия горутин, у них есть общие аспекты:
	1) внутри горутины берут строку из канала, выводят ее в консоль и кладут обратно в канал
	2) получают сигнал о необходимости закрытия перед концом программы
	3) везде используется sync.WaitGroup, для того, чтобы программа дождалась правильного завершения горутины*/

	// Функция signalTerm получает сигнал Ctrl+C, который передается в горутину, для того, чтобы она узнала, что нужно закрываться
	signalTerm()
	// Функция withDeadLine получает сигнал о закрытии, спустя определенное время, после начала программы
	withDeadLine()
	// Функция withTimeout получает сигнал о закрытии, спустя определенное время, после начала программы
	withTimeout()
	// Функция withCancel получает сигнал ошибки, и по этому отправляет сигнал открывшимся горутинам, что пора закрываться
	withCancel()
}

func withDeadLine() {
	wg := sync.WaitGroup{}

	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	ch := make(chan string, 1)
	ch <- "Hello World"

	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			select {
			case <-ctx.Done():
				fmt.Println("Close go func")
				close(ch)
				return
			default:
				mes := <-ch
				fmt.Println(mes)
				ch <- mes
			}
		}
	}()

	wg.Wait()
}

func withTimeout() {
	wg := sync.WaitGroup{}

	ctx := context.Background()
	d := time.Now().Add(3 * time.Second)
	ctx, cancel := context.WithDeadline(ctx, d)
	defer cancel()

	ch := make(chan string, 1)
	ch <- "Hello World"

	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			select {
			case <-ctx.Done():
				fmt.Println("Close go func")
				close(ch)
				return
			default:
				mes := <-ch
				fmt.Println(mes)
				ch <- mes
			}
		}
	}()

	wg.Wait()
}

func longRunningProcess() error {
	time.Sleep(time.Second)
	return errors.New("Failed: Process took too long time")
}

func withCancel() {
	wg := sync.WaitGroup{}

	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)

	go func() {
		err := longRunningProcess()
		if err != nil {
			cancel()
		}
	}()

	ch := make(chan string, 1)
	ch <- "Hello World"

	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			select {
			case <-ctx.Done():
				fmt.Println("Close go func")
				close(ch)
				return
			default:
				mes := <-ch
				fmt.Println(mes)
				ch <- mes
			}
		}
	}()

	wg.Wait()
}

func signalTerm() {
	wg := sync.WaitGroup{}

	sigs := make(chan os.Signal, 1)                      // Открываем канал в который будет подаваться сигнал о завершении программы
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM) // говорим каналу sigs, чтобы он ждал только сигнал завершения программы

	ch := make(chan string, 1)
	ch <- "Hello World"

	wg.Add(1)

	go func() {
		defer wg.Done()
		for {
			select {
			case <-sigs:
				fmt.Println("Close go func")
				close(ch)
				return
			default:
				mes := <-ch
				fmt.Println(mes)
				ch <- mes
			}
		}
	}()

	wg.Wait()
}
