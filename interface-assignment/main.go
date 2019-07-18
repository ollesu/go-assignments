package main

import "fmt"

// Both have func getArea that returns calculated area of the square or triangle

// area of triangle = 0.5 * base * height
// square - sideLength * sideLength

// add interface "shape" that defines func printArea
// - calculate the area of the given shape
// - func can be called either with square or triangle

type square struct {
	sideLength float64
}

type triangle struct {
	height float64
	base   float64
}

type shape interface {
	getArea() float64
}

func main() {
	tr := triangle{
		base:   20,
		height: 20,
	}

	s := square{
		sideLength: 30,
	}

	printArea(tr)
	printArea(s)
}

func (s square) getArea() float64 {
	return s.sideLength * s.sideLength
}

func (t triangle) getArea() float64 {
	return 0.5 * t.base * t.height
}

func printArea(s shape) {
	fmt.Println("The calculated area is", s.getArea())

}
