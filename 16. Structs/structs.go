// Structures in Go are typed collections of fields
package main

import (
	"fmt"
)

// Declaring a "person" struct
type person struct {
	name string
	age  int
}

func main() {
	// Create new struct
	fmt.Println(person{"Alpha", 20})

	// Create new struct with key index
	fmt.Println(person{name: "Beta", age: 30})

	// Omitted fields will be zero-valued
	fmt.Println(person{name: "Gamma"})

	// & prefix yields a pointer to the struct
	fmt.Println(&person{name: "Delta", age: 40})

	// Access a struct's field with a dot
	s := person{name: "Epsilon", age: 50}
	fmt.Println(s.name)

	// Creating a structs pointer and using dots
	// Note: the pointers are automatically dereferenced!
	sp := &s
	fmt.Println(sp.age)

	// Structs are mutable
	sp.age = 51
	fmt.Println(sp.age)
}
