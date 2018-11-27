package main

import (
	"fmt"
	"os"
)

func main() {
	// defer(s) will not be run when using os.Exit,
	// so this fmt.Println will never be called.
	defer fmt.Println("!")

	// Exit with status 3
	os.Exit(3)
}
