package main

import "fmt"

func main() {

	var a string = "initial"
	fmt.Println(a)

	var b, c int = 1, 2
	fmt.Println(b, c)

	var d bool = true
	fmt.Println(d)

	var e int // declare w/o a corresponding initialization are zero-valued
	fmt.Println(e)

	f := "short" // Shorthand for var f = "short"
	fmt.Println(f)

	var g, h, i = true, false, "no!"
	fmt.Println(g, h, i)
}
