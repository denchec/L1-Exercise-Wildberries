package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var numSec time.Duration
	_, err := fmt.Scan(&numSec) // Считываем с клавиатуры количество секунд, через сколько выключить программу
	checkError(err)

	wg := sync.WaitGroup{} // Присваиваем вспомогательную функцию, для ожидания завершения горутин, переменной wg

	ch := make(chan int, 1) // Создаем новый канал

	exitTime := time.After(time.Second * numSec) // Говорим программе, через сколько завершиться

	sequenceOfNumbers := []int{}

	wg.Add(1) // Добавляем 1 к количеству горутин, завершения которых должен будет ждать wg.Wait()
	go func() {
		num := 1
		for {
			select {
			// Если время истекло, то: wg.Done() отнимет 1 из счетчика незакрытых горутин и return завершит горутину
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

	wg.Wait()                           // Функция wg.Wait() не дает завершиться горутине main. Как только все горутины завершатся, main будет закрыт
	fmt.Println(len(sequenceOfNumbers)) // Выводим количество новых записей, созданных за отведенный промежуток времени
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
