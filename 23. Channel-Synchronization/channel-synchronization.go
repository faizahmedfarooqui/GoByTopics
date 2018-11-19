/**
 * We can use channels to synchronize execution across go-routines.
 * Here, we have an example of using a blocking receive to wait for
 * a go-routine to finish.
 */
package main

import (
	"fmt"
	"time"
)

func worker(done chan bool) {
	fmt.Println("working...")
	time.Sleep(time.Second)
	fmt.Println("done")

	done <- true
}

func main() {
	done := make(chan bool, 1)
	go worker(done)

	// Block until we receive a notification from the worker on the channel.
	// If we remove the below line, the program will exit before the
	// "worker" even started.
	<-done
}
