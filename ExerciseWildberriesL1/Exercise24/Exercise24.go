package main

import (
	"fmt"
	"math"
)

type Point struct { // Создаем структуру Point со значениями x, y
	x float64
	y float64
}

func main() {
	/* Используя функцию constructor, передаем ее аргументы переменным aPoint и bPoint,
	которые будут иметь структуру Point */
	aPoint := constructor(1, 2)
	bPoint := constructor(3, 4)
	/* Используем функцию distance, обращаемся к встроенной библиотеке math, и с помощью нее
	высчитываем расстояние между точек*/
	fmt.Println(distance(aPoint, bPoint))
}

func distance(a *Point, b *Point) float64 {
	return math.Sqrt(math.Pow(a.x-b.x, 2) + math.Pow(a.y-b.y, 2))
}

func constructor(x, y float64) *Point {
	return &Point{x: x, y: y}
}
