package main

import "fmt"

type Human struct { // Определили структуру Human
	Name   string
	Age    int
	Gender string
}

type Action struct { // Определили структуру Action, которая дополнительно ссылается на структуру Human
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

	// Теперь мы можем вывести данные переменной на экран, либо обратившись к переменной
	fmt.Println(FirstPerson)

	// Либо обратившись к переменной с определенным полем
	fmt.Println(FirstPerson.Name, FirstPerson.Age, FirstPerson.Height)
}
