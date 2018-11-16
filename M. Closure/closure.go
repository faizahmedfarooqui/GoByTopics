package main

import (
	"fmt"
)

// This function returns another function
// Note: The returned function closes over
// the variable i to form a closure
func intSeq() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

func main() {
	// function captures its own "i" value,
	// which will be updated on its every call
	nextInt := intSeq()

	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())

	// On every call, the state is unique to that
	// particular function, create/test a new one.
	newInts := intSeq()
	fmt.Println(newInts())
}
