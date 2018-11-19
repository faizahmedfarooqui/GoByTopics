/**
 * Select lets you wait on multiple channel operations.
 *
 * Combining "go-routines" and "channels" with select is
 * a powerful feature of "Go".
 */
package main

import (
	"fmt"
	"time"
)

func main() {

	c1 := make(chan string)
	c2 := make(chan string)

	go func() {
		time.Sleep(1 * time.Second)
		c1 <- "one"
	}()

	go func() {
		time.Sleep(2 * time.Second)
		c2 <- "two"
	}()

	// Here we are using select to await both of the values
	// simultaneously, printing each one as it is arrives
	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-c1:
			fmt.Println("received:", msg1)
		case msg2 := <-c2:
			fmt.Println("received:", msg2)
		}
	}

	/**
	 * Try running this program using "time go run select.go"
	 *
	 * You will notice that real time is ~2 seconds since both
	 * sleep execute concurrently.
	 */
}
