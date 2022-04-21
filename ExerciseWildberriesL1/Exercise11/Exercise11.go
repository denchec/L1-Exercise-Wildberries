package main

import "fmt"

type void struct{}

/* Объявляем тип set, у которого вместо значения пустая структура(пустая структура не занимает место в памяти)
Передав этот тип переменной, мы автоматически делаем ее множеством*/
type set map[int]void

// Объявляем общую переменную с пустой структурой, которую мы будем использовать как значение для множества
var member void

func newSet(arr ...int) set {
	s := set{}
	for _, i := range arr { // Берем аргументы по 1 и добавляем в множество
		s[i] = member
	}
	return s
}

func (s set) Intersect(s2 set) set {
	intersection := set{}
	for key, _ := range s {
		_, keyContainsInS2 := s2[key] // Проверяем, есть ли ключ из первого множества, во втором
		if keyContainsInS2 {          // Если есть - выписываем его в отдельное множество
			intersection[key] = member
		}
	}
	return intersection
}

func main() {
	s1 := newSet(1, 2, 3, 4, 5)
	s2 := newSet(7, 15, 3, 22, 5)

	fmt.Println(s1.Intersect(s2)) // Выводим множество пересекающихся значений
}
