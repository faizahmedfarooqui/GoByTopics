package main

import "fmt"

func main() {

	i := 1
	for i <= 3 {
		fmt.Println(i)
		i++ // increment syntax or you can replace that with i += 1
	}

	for j := 7; j <= 9; j++ {
		fmt.Println(j)
	}

	for {
		fmt.Println("loop")
		break // break keyword
	}

	for n := 0; n <= 5; n++ {
		if n%2 == 0 {
			continue // continue keyword
		}

		fmt.Println(n)
	}
}
