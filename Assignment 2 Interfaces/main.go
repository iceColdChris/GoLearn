package main

import (
	"fmt"
)

type shape interface {
	getArea() float64
}

type triangle struct {
	height float64
	base   float64
}
type square struct {
	sideLength float64
}

func main() {
	sq := square{5}
	printArea(sq)

	tr := triangle{base: 5, height: 5}
	printArea(tr)
}

func printArea(s shape) {
	a := s.getArea()
	fmt.Println("The area is", a)
}

func (s square) getArea() float64 {
	return s.sideLength * s.sideLength
}

func (t triangle) getArea() float64 {
	return 0.5 * t.base * t.height
}
