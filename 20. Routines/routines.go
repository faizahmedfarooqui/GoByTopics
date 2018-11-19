/**
 * Routine is a lightweight thread managed by the "Go" runtime
 */
package main

import (
	"fmt"
)

func f(from string) {
	for i := 0; i < 3; i++ {
		fmt.Println(from, ":", i)
	}
}

func main() {

	// This is usual function call
	// and this runs synchronously
	f("direct")

	// To invoke a function in go's routine
	// use "go" keyword before calling the
	// function
	go f("this")

	// Running a function in async
	go func(msg string) {
		fmt.Println(msg)
	}("going")

	// Ask for user's to press any key
	// Note: Because of this the program is
	// waiting meanwhile we'll get async
	// call's output
	fmt.Scanln()

	fmt.Println("done")
}
