package main

import (
	"crypto/sha1"
	"fmt"
)

func main() {
	s := "sha1 this string"

	// The pattern for generating a hash is sha1.New(),
	// sha1.Write(bytes), then sha1.Sum([]byte{}).

	// To compute MD5 hashes import crypto/md5 & use md5.New()

	h := sha1.New()

	h.Write([]byte(s))

	bs := h.Sum(nil)
	fmt.Println(s)

	// SHA1 values are often printed in hex, for example
	// in git commits. Use the %x format verb to convert
	// a hash results to a hex string.
	fmt.Printf("%x\n", bs)
}
