// Pointers in Go allows you to pass references
// to values & records within the program
package main

import (
	"fmt"
)

// this function does not change "i"
// in "main" function
func zeroVal(ival int) {
	ival = 0
}

// this function will change "i"
// in "main" function
func zeroPtr(iptr *int) {
	*iptr = 0
}

func main() {
	i := 1
	fmt.Println("initial: ", i)

	zeroVal(i)
	fmt.Println("zeroVal:", i)

	// "&i" syntax gives the memory address of "i"
	// ie. a pointer to "i"
	zeroPtr(&i)
	fmt.Println("zeroPtr:", i)

	// Pointers can be printed too...!
	fmt.Println("pointer:", &i)
}
