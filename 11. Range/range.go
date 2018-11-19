// Range iterates over elements in a variety of data structures
package main

import (
	"fmt"
)

func main() {

	nums := []int{2, 3, 4}
	sum := 0

	// Sum the numbers array in Slice.
	// Note: Here, we dont need index so we
	// ignored index by blank identifier.
	for _, num := range nums {
		sum += num
	}
	fmt.Println("sum:", sum)

	// Slice array with the index
	for i, num := range nums {
		if num == 3 {
			fmt.Println("index:", i)
		}
	}

	// Range on map over key-value pair
	kvs := map[string]string{"a": "apple", "b": "banana"}
	for k, v := range kvs {
		fmt.Printf("%s -> %s\n", k, v)
	}

	// Range can iterate over just keys of a map
	for k := range kvs {
		fmt.Println("key:", k)
	}

	// Range over strings iterate over Unicode
	// code points. First value is starting byte
	// index of the rune and the second the rune itself
	for i, c := range "Go" {
		fmt.Println(i, c)
	}
}
