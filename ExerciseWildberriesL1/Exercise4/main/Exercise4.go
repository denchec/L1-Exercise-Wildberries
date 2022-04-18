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
	_, err := fmt.Scan(&amountOfWorkers) // Вводим с консоли необходимое количество воркеров
	checkError(err)

	var wg sync.WaitGroup // Присваиваем вспомогательную функцию, для ожидания завершения горутин, переменной wg

	channel := make(chan string, amountOfWorkers) // Открываем канал с тем же количеством слотов, сколько и воркеров

	sigs := make(chan os.Signal, 1)                      // Открываем канал в который будет подаваться сигнал о завершении программы
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM) // говорим каналу sigs, чтобы он ждал только сигнал завершения программы

	// Создаем цикл, в котором будет генерироваться столько разных строк, сколько воркеров, а так же отправляем их в канал channel
	for i := 1; i < amountOfWorkers+1; i++ {
		newWord := strings.Repeat("x", i) + "HELLO" + strings.Repeat("x", i)
		channel <- newWord
		wg.Add(1) // Добавляем 1 к количеству горутин, завершения которых должен будет ждать wg.Wait()
		go func() {
			defer wg.Done() // Говорим функции wg.Wait(), что 1 горутина завершена(-1 к общему счетчику)
			for {
				select {
				/* Если сигнал о завершении программы будет получен, то:
				1) мы запишем данные сигнала в переменную signalSigterm
				2) вернем сигнал из переменной signalSigterm обратно в канал, чтобы другие горутины могли его использовать
				3) return завершит горутину*/
				case signalSigterm := <-sigs:
					sigs <- signalSigterm
					return
					// Если сигнал о завершении программы не получен, то: получаем строку из канала, выводим ее и отправляем обратно в канал
				default:
					sayHello := <-channel
					_, err := fmt.Fprintln(os.Stdout, sayHello)
					checkError(err)
					channel <- sayHello
				}
			}
		}()
	}

	wg.Wait() // Функция wg.Wait() не дает завершиться горутине main. Как только все горутины завершатся, main будет закрыт
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
