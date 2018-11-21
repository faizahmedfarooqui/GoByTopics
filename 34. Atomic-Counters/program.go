/**
 * Primary mechanism of managing state in Go is communication over channels.
 * And it was done in Worker-Pools topic, but there are few other options
 * for managing state...!
 *
 * Here we will look at using, sync/atomic package for atomic counters
 * accessed by multiple go-routine.
 */
package main

import (
	"fmt"
	"sync/atomic"
	"time"
)

func main() {
	// "Always-positive" integer as counter
	var ops uint64

	// Running 50 go-routines to increment the counter.
	// Note: "atomic.AddUint64" increments the counter
	// with the "&" syntax!
	for i := 0; i < 50; i++ {
		go func() {
			for {
				atomic.AddUint64(&ops, 1)

				time.Sleep(time.Millisecond)
			}
		}()
	}

	time.Sleep(time.Second)

	// We are safely going to extract the value while its
	// might still be getting updated by some go routines.
	// We use "atomic.LoadUint64" over "ops" address
	// using "&" syntax!
	opsFinal := atomic.LoadUint64(&ops)

	// This output will show the number of operations we
	// were able to perform!
	fmt.Println("ops:", opsFinal)
}
