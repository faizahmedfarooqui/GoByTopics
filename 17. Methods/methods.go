// In Go, methods are defined on struct types
package main

import (
	"fmt"
)

type rect struct {
	width, height int
}

// "area" method has "pointer receiver" type of *rect
func (r *rect) area() int {
	return r.width * r.height
}

// "perim" method has "value receiver" type of rect
func (r rect) perim() int {
	return 2 * (r.width + r.height)
}

func main() {
	r := rect{width: 10, height: 5}

	fmt.Println("Area:", r.area())
	fmt.Println("Perimeter:", r.perim())

	rp := &r

	fmt.Println("Area:", rp.area())
	fmt.Println("Perimeter:", rp.perim())
}
