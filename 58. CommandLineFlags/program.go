/**
 * What are CL-flags?
 * Ans: in wc -l the -l is a command-line flag.
 */
package main

import (
	"flag"
	"fmt"
)

func main() {
	// Basic flag declarations are available for string, integer,
	// and boolean options. Here we declare a string flag word
	// with a default value "foo" and a short description. This
	// flag.String function returns a string pointer (not a string
	// value); we’ll see how to use this pointer below.
	wordPtr := flag.String("word", "foo", "a string")

	// This declares numb and fork flags, using a similar approach
	// to the word flag.
	numbPtr := flag.Int("numb", 42, "an int")
	boolPtr := flag.Bool("fork", false, "a bool")

	// It’s also possible to declare an option that uses an existing
	// var declared elsewhere in the program. Note that we need to
	// pass in a pointer to the flag declaration function.
	var svar string
	flag.StringVar(&svar, "svar", "bar", "a string var")

	// Once all flags are declared, call flag.Parse() to execute the
	// command-line parsing.
	flag.Parse()

	fmt.Println("word:", *wordPtr)
	fmt.Println("numb:", *numbPtr)
	fmt.Println("fork:", *boolPtr)
	fmt.Println("svar:", svar)
	fmt.Println("tail:", flag.Args())

	/**
	 * To run, build using command "go build filename.go"
	 * Then run, "./filename -word=opt -numb=7 -fork -svar=flag"
	 * Also try, "./filename"
	 * And for help, "./filename -h"
	 */
}
