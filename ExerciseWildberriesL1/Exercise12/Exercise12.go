package main

import "fmt"

type void struct{}

type set map[string]struct{}

var member void

func main() {
	newSlice := []string{"cat", "cat", "dog", "cat", "tree"}
	s := set{} // Объявляем переменную типа set

	// Записываем в map все значения среза, тем самым делая из значений среза, множество
	for i := 0; i < len(newSlice); i++ {
		s[newSlice[i]] = member
	}

	fmt.Println(s)
}
