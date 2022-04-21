package main

import (
	"fmt"
	"sync"
)

type Counter struct { // Создаем структуру, в которой будет находиться наш счетчик
	FinishCount int
}

func main() {
	count := Counter{} // Передаем структуру Counter переменной

	wg := sync.WaitGroup{} // Передаем функцию WaitGroup переменной

	for i := 0; i < 100; i++ { // Запускаем 100 горутин
		wg.Add(1) // Добавляем 1 к счетчику ожидания завершения горутины
		go func() {
			defer wg.Done()                // При завершении горутины отнимаем 1 из счетчика ожидания завершения
			count.FinishCount++            // Инкриминируем счетчик
			fmt.Println(count.FinishCount) // Выводим счетчик для отслеживания конкуренции
		}()
	}

	wg.Wait() // Ждем завершения всех горутин

	fmt.Println(count.FinishCount) // Выводим общее значение счетчика
}
