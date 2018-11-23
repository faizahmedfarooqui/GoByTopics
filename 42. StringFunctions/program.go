/**
 * The standard library’s strings package provides many useful
 * string-related functions.
 *
 * There a lot of them, you can dig more into it from:
 * https://golang.org/pkg/strings
 * https://blog.golang.org/strings
 */
package main

import (
	"fmt"
	s "strings"
)

// Creating a shorthand for the long syntax
var p = fmt.Println

func main() {
	p("Contains:  ", s.Contains("test", "es"))
	p("Count:     ", s.Count("test", "t"))
	p("HasPrefix: ", s.HasPrefix("test", "te"))
	p("HasSuffix: ", s.HasSuffix("test", "st"))
	p("Index:     ", s.Index("test", "e"))
	p("Join:      ", s.Join([]string{"a", "b"}, "-"))
	p("Repeat:    ", s.Repeat("a", 5))
	p("Replace:   ", s.Replace("foo", "o", "0", -1))
	p("Replace:   ", s.Replace("foo", "o", "0", 1))
	p("Split:     ", s.Split("a-b-c-d-e", "-"))
	p("ToLower:   ", s.ToLower("TEST"))
	p("ToUpper:   ", s.ToUpper("test"))
	p()

	p("Len: ", len("hello")) // Get string length
	p("Char:", "hello"[1])   // Get byte of the char
}
