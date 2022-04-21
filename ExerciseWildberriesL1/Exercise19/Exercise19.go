package main

import "fmt"

func main() {
	var messageWord string
	_, err := fmt.Scan(&messageWord)
	checkError(err)

	var newWord string

	for _, i := range messageWord {
		newWord = string(i) + newWord
	}

	fmt.Println(newWord)
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
