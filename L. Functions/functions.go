package main

import (
	"fmt"
)

// Single Explicit Return
func plus(a int, b int, c int) int {
	return a + b + c
}

// Multiple Explicit Return
func vals() (int, int) {
	return 3, 7
}

// Multiple Explicit Return with shared type
func swap(x, y string) (string, string) {
	return y, x
}

// Named return values
// Note: A "return" statement w/o arguments
// returns the named return values
func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

// Variadic functions can be called with any
// number of arguments. For ex. "fmt.Println"
func sum(nums ...int) (total int) {
	fmt.Println(nums, " ")

	total = 0
	for _, num := range nums {
		total += num
	}

	return
}

func main() {

	// 1. It is mandatory in go that you should provide explicit return!

	fmt.Println("Single Return:", plus(1, 2, 3))

	a, b := vals()
	c, _ := vals()
	_, d := vals()

	// 2. We cannot pass "vals()" inside "Println"
	// because "Println" is single-value context
	// and "vals()" return multiple values

	// 3. We can get a subset of the returned values,
	// then use blank identifier

	fmt.Println("Multiple Return into two variables:", a, b)
	fmt.Println("Multiple Return: c = ", c)
	fmt.Println("Multiple Return: d = ", d)

	// 4. Multiple Explicit Return with shared type
	e, f := swap("hello", "world")
	fmt.Println("Multiple with Shared-Type Return:", e, f)

	// 5. Named return value
	g, h := split(17)
	fmt.Println("Named Return:", g, h)

	// 6. Variadic
	nums := []int{1, 2, 3, 4}
	fmt.Println("Variadic return:", sum(nums...))
}
