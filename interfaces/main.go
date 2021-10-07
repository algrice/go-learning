package main

import "fmt"

type triangle struct {
	baseLength float64
	height     float64
}

type square struct {
	sideLength float64
}

func (t triangle) getArea() float64 {
	return (t.baseLength * t.height) / 2
}

func (s square) getArea() float64 {
	return s.sideLength * s.sideLength
}

type shape interface {
	getArea() float64
}

func printArea(s shape) {
	fmt.Println(s.getArea())
}

func main() {
	testTriangle := triangle{3, 3}
	testSquare := square{10}
	printArea(testTriangle)
	printArea(testSquare)
}
