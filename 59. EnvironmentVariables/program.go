package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	// To set a key/value pair, use os.Setenv
	os.Setenv("FOO", "BAR")

	// To get a value for a key, use os.Getenv.
	// This will return an empty string if the key isn’t present in the environment.
	fmt.Println("FOO:", os.Getenv("FOO"))

	fmt.Println("BAR:", os.Getenv("BAR"))

	fmt.Println()

	// Use os.Environ to list all key/value pairs in the environment.
	// This returns a slice of strings in the form KEY=value.
	for _, e := range os.Environ() {
		pair := strings.Split(e, "=")
		fmt.Println(pair[0])
	}
}
