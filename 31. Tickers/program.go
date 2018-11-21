/**
 * Tickers are for when you want to do something repeatedly at regular intervals.
 */
package main

import (
	"fmt"
	"time"
)

func main() {
	ticker := time.NewTicker(500 * time.Millisecond)
	go func() {
		for t := range ticker.C {
			fmt.Println("Tick at", t)
		}
	}()

	time.Sleep(1600 * time.Millisecond)

	// Ticker can be stopped same as Timers
	ticker.Stop()
	fmt.Println("Ticker stopped")
}
