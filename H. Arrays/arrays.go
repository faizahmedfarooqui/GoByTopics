package main

import "fmt"

func main() {

	// By default arrays in Go are zero-valued
	var a [5]int
	fmt.Println("emp:", a)

	a[4] = 100
	fmt.Println("set:", a)
	fmt.Println("get:", a[4])

	// Built-in "len" to get array length
	fmt.Println("len:", len(a))

	// Syntax to declare & initialize an array
	// in one line
	b := [5]int{1, 2, 3, 4, 5}
	fmt.Println("dcl:", b)

	// Array types are one-D, but we can compose
	// types to build multi-D structures.
	var twoD [2][3]int
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			twoD[i][j] = i + j
		}
	}

	fmt.Println("2d:", twoD)
}
