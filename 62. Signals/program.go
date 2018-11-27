/**
 * Go programs can intelligently handle Unix signals.
 */
package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	// Go signal notification works by sending os.Signal values on a
	// channel. We’ll create a channel to receive these notifications
	// (we’ll also make one to notify us when the program can exit).
	sigs := make(chan os.Signal, 1)
	done := make(chan bool, 1)

	// signal.Notify registers the given channel to receive
	// notifications of the specified signals.
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	// This goroutine executes a blocking receive for signals. When it
	// gets one it’ll print it out and then notify the program that it
	// can finish.
	go func() {
		sig := <-sigs
		fmt.Println()
		fmt.Println(sig)
		done <- true
	}()

	// Wait here
	fmt.Println("awaiting signals")
	<-done
	fmt.Println("exiting")
}
