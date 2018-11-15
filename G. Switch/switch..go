package main

import (
	"fmt"
	"time"
)

func main() {

	i := 2
	fmt.Println("Write", i, "as")

	// Basic with use of "default" keyword
	switch i {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	default:
		fmt.Println("something else")
	}

	// Can use commas to separate multiple
	// expressions in same case statement
	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("It's the Weekend!")
	default:
		fmt.Println("It's a Weekday")
	}

	// Switch w/o expression is an alternate
	// way to express if/else. Here we show how
	// a case statement can be non-constant.
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("It's before noon")
	default:
		fmt.Println("It's after noon")
	}

	// A "type" switch compares types instead of
	// values. We can use it to discover the type
	// of an interface value. Note: "%T" in Printf
	// will print the type of variable.
	whatAmI := func(i interface{}) {
		switch t := i.(type) {
		case bool:
			fmt.Println("I'm a boolean")
		case int:
			fmt.Println("I'm an integer")
		default:
			fmt.Printf("Don't know type %T\n", t)
		}
	}

	whatAmI(true)
	whatAmI(1)
	whatAmI("hey")
}
