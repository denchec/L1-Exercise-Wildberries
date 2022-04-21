package main

import "fmt"

type Human struct {
	Name   string
	Age    int
	Gender string
}

type Action struct {
	Height float32
	Weight float32
	Human
}

func main() {
	// Вложили в переменную FirstPerson информацию по человеку с именем Denis
	FirstPerson := Action{
		Height: 171,
		Weight: 60,
		Human: Human{
			Name:   "Denis",
			Age:    25,
			Gender: "Man"}}

	// Выводим данные переменной на экран обратившись к переменной
	fmt.Println(FirstPerson)

	// Выводим данные переменной на экран обратившись к переменной с определенным полем
	fmt.Println(FirstPerson.Name, FirstPerson.Age, FirstPerson.Height)
}
