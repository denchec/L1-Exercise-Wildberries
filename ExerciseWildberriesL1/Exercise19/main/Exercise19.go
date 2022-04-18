package main

import "fmt"

func main() {
	var messageString string
	_, err := fmt.Scan(&messageString)
	checkError(err)

	var newString string

	for _, i := range messageString {
		newString = string(i) + newString
	}

	fmt.Println(newString)
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
