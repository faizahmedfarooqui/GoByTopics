/**
 * Timers represents a single event in the future.
 */
package main

import (
	"fmt"
	"time"
)

func main() {
	// This timer will wait for 2 seconds
	timer1 := time.NewTimer(2 * time.Second)

	// This line blocks on the timer's channel "C" until
	// it sends a value indicating that the timer expired
	<-timer1.C
	fmt.Println("Timer 1 expired")

	// If you just wanted to wait, you could have used
	// "time.Sleep". One reason a timer may be useful
	// is that you can cancel the timer before it expires.
	timer2 := time.NewTimer(time.Second)
	go func() {
		<-timer2.C
		fmt.Println("Timer 2 expired")
	}()

	stop2 := timer2.Stop()
	if stop2 {
		fmt.Println("Timer 2 stopped")
	}
}
