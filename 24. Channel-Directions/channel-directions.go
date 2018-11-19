/**
 * When using channel as function parameters, you can
 * specify if a channel is meant to only send or receive values.
 *
 * This specificity increases the type-safety of the program.
 */
package main

import (
	"fmt"
)

func ping(pings chan<- string, msg string) {
	pings <- msg
}

func pong(pings <-chan string, pongs chan<- string) {
	msg := <-pings
	pongs <- msg
}

func main() {
	pings := make(chan string, 1)
	pongs := make(chan string, 1)

	ping(pings, "passed a message")
	pong(pings, pongs)

	fmt.Println(<-pongs)
}
