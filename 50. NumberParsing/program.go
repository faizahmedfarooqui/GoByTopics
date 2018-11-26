package main

import (
	"fmt"
	"strconv"
)

func main() {
	// 64 here, tells how many bits of precision to parse
	f, _ := strconv.ParseFloat("1.234", 64)
	fmt.Println(f)

	// 0 here, means infer the base from the string
	// 64 required that result fit in 64 bits
	i, _ := strconv.ParseInt("123", 0, 64)
	fmt.Println(i)

	// ParseInt will recognize hex-formatted numbers
	d, _ := strconv.ParseInt("0x1c8", 0, 64)
	fmt.Println(d)

	u, _ := strconv.ParseUint("789", 0, 64)
	fmt.Println(u)

	// Atoi is a convenience function for basic base-10
	// int parsing
	k, _ := strconv.Atoi("135")
	fmt.Println(k)

	// Catch error on bad input
	_, e := strconv.Atoi("wat")
	fmt.Println(e)
}
