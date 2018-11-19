package main

import "fmt"

func main() {

	// Can have "if" statement w/o "else"
	// & w/o "paranthesis"
	if 8%4 == 0 {
		fmt.Println("8 is divisible by 4")
	}

	// If-Else "ladder"
	if 7%2 == 0 {
		fmt.Println("7 is even")
	} else {
		fmt.Println("7 is odd")
	}

	// Can start with the condition; any variables
	// declared in this statement will be available
	// in all branches.
	// Note: There is no ternary if in "Go" (ie. ?:)
	if num := 9; num < 0 {
		fmt.Println(num, "is negative")
	} else if num < 10 {
		fmt.Println(num, "has 1 digit")
	} else {
		fmt.Println(num, "has multiple digits")
	}
}
