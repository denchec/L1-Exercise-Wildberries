package main

import "fmt"

func main() {
	var A int
	var B int

	_, aErr := fmt.Scan(&A)
	checkError(aErr)
	_, bErr := fmt.Scan(&B)
	checkError(bErr)

	C := A * B
	fmt.Println("A * B", C)
	C = A + B
	fmt.Println("A + B", C)
	C = A - B
	fmt.Println("A - B", C)
	C = A % B
	fmt.Println("A % B", C)
	C = A / B
	fmt.Println("A / B", C)
	var D float64
	D = float64(A) / float64(B)
	fmt.Println("A / B", D)
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
