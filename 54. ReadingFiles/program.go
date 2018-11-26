package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	// the most basic file reading task is slurping a file’s entire contents into memory.
	dat, err := ioutil.ReadFile("./dat")
	check(err)
	fmt.Println(string(dat))

	// You’ll often want more control over how and what parts of a file are read.
	// For these tasks, start by Opening a file to obtain an os.File value.
	f, err := os.Open("./dat")
	// defer d.Close() // You can do this or f.Close() at the end of the main function
	check(err)

	// Read some bytes from the beginning of the file. Allow up to 5 to be read but
	// also note how many actually were read.
	b1 := make([]byte, 5)
	n1, err := f.Read(b1)
	check(err)
	fmt.Printf("%d bytes: %s\n", n1, string(b1))

	// You can also Seek to a known location in the file and Read from there.
	o2, err := f.Seek(6, 0)
	check(err)
	b2 := make([]byte, 2)
	n2, err := f.Read(b2)
	check(err)
	fmt.Printf("%d bytes @ %d: %s\n", n2, o2, string(b2))

	// The io package provides some functions that may be helpful for file reading.
	// For example, reads like the ones above can be more robustly implemented with ReadAtLeast.
	o3, err := f.Seek(6, 0)
	check(err)
	b3 := make([]byte, 2)
	n3, err := io.ReadAtLeast(f, b3, 2)
	check(err)
	fmt.Printf("%d bytes @ %d: %s\n", n3, o3, string(b3))

	// There is no built-in rewind, but Seek(0, 0) accomplishes this.
	_, err = f.Seek(0, 0)
	check(err)

	// The bufio package implements a buffered reader that may be useful both for its efficiency
	// with many small reads and because of the additional reading methods it provides.
	r4 := bufio.NewReader(f)
	b4, err := r4.Peek(5)
	check(err)
	fmt.Printf("5 bytes: %s\n", string(b4))

	// Close the file when you’re done (usually this would be scheduled immediately after Opening with defer).
	f.Close()
}
