package main

import (
	"fmt"
	"sort"
)

func main() {
	// String type slice
	strs := []string{"c", "a", "b"}
	sort.Strings(strs)
	fmt.Println("Strings:", strs)

	// Integer type slice
	ints := []int{7, 2, 4}
	sort.Ints(ints)
	fmt.Println("Integers:", ints)

	// Check if the slice is already sorted
	s := sort.IntsAreSorted(ints)
	fmt.Println("Sorted:", s)
}
