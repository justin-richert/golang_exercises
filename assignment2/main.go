package main

import "fmt"

type shape interface {
	getArea() float64
}

type square struct {
	sideLength float64
}

type triangle struct {
	base   float64
	height float64
}

func main() {
	t := triangle{
		base:   15,
		height: 12,
	}
	s := square{
		sideLength: 9,
	}
	printArea(t)
	printArea(s)
}

func printArea(sh shape) {
	fmt.Println("Area:", sh.getArea())
}

func (s square) getArea() float64 {
	return s.sideLength * s.sideLength
}

func (t triangle) getArea() float64 {
	return 0.5 * t.base * t.height
}
