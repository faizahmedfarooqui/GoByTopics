/**
 * Panic means something went unexpectedly wrong.
 *
 * Mostly we use it to fail fast on errors that shouldn't
 * occur during normal operation, or that we aren't prepared
 * to handle gracefully.
 */
package main

import (
	"os"
)

func main() {
	panic("a problem")

	// Handle if file creation fails due to unexpected reasons
	_, err := os.Create("/tmp/file")
	if err != nil {
		panic(err)
	}

	/**
	 * Running this program will cause it to panic, print an
	 * error message and goroutine traces, and exit with a
	 * non-zero status.
	 *
	 * Note that unlike some languages which use exceptions
	 * for handling of many errors, in Go it is idiomatic to
	 * use error-indicating return values wherever possible.
	 */
}
