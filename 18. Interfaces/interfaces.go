package main

import (
	"fmt"
	"math"
)

// A Geometry interface with 2 methods
type geometry interface {
	area() float64
	perim() float64
}

// Rectangle struct with 2 properties
type rect struct {
	width, height float64
}

// Circle struct with 1 property
type circle struct {
	radius float64
}

// area method for rect structure
func (r rect) area() float64 {
	return r.width * r.height
}

// perim method for rect structure
func (r rect) perim() float64 {
	return 2 * (r.width + r.height)
}

// area method for circle structure
func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

// perim method for circle structure
func (c circle) perim() float64 {
	return 2 * math.Pi * c.radius
}

// Implementation of geometry interface
func measure(g geometry) {
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.perim())
}

func main() {
	r := rect{width: 3, height: 4}
	c := circle{radius: 5}

	measure(r)
	measure(c)
}
