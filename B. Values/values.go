package main

import "fmt"

func main() {

	// Default values;
	// 1. 0 for numeric
	// 2. false for boolean
	// 3. "" (empty) for strings

	fmt.Println("go" + "lang")

	fmt.Println("1+1 =", 1+1)
	fmt.Println("7.0/3.0 =", 7.0/3.0)
	fmt.Println(true && false)
	fmt.Println(true || false)
	fmt.Println(!true)
}
