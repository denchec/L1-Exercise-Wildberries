package main

func main() {
	a := 1
	b := 2
	println("a =", a, "b =", b)

	a, b = b, a

	println("a =", a, "b =", b)
}
