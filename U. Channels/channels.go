/**
 * Channels are the pipes that connect concurrent goroutines. '
 * You can send values into channels from one goroutine and
 * receive those values into another goroutine.
 */
package main

import (
	"fmt"
)

func main() {
	// Syntax for creating new channel: "make(chan val-type)"
	// Channels are typed by the values they convey
	messages := make(chan string)

	// Send a value into a channe; using the channel "<-" snytax
	go func() { messages <- "ping" }()

	// Receive values using "<-" syntax from the created channel
	msg := <-messages

	// By default, sends & receives block until the other side
	// is ready. This allows goroutines to synchronize without
	// explicit locks or condition variables.
	fmt.Println(msg)

	// Now, an example
	s := []int{7, 2, 8, -9, 4, 0}

	c := make(chan int)

	go sum(s[:len(s)/2], c)
	go sum(s[len(s)/2:], c)

	// Receive from "c" channel
	x, y := <-c, <-c

	fmt.Println(x, y, x+y)
}

func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}

	// send "sum" to "c" channel
	c <- sum
}
