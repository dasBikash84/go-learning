package main

import "fmt"

type HasArea interface {
	printArea()
}

type triangle struct {
	height float64
	base   float64
}
type square struct {
	sideLength float64
}

func (t triangle) printArea() {
	fmt.Println(0.5 * t.base * t.height)
}

func (s square) printArea() {
	fmt.Println(s.sideLength * s.sideLength)
}

func main() {

	t := triangle{2, 4}
	s := square{2}

	t.printArea()
	s.printArea()

}
