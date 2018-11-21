/**
 * Primary mechanism of managing state in Go is communication over channels.
 * And it was done in Worker-Pools topic, Atomic-Counters, but there are
 * few more options for managing state...!
 *
 * Here we will look at using, "mutex" package to safely access data across
 * multiple go-routines.
 */
package main

import (
	"fmt"
	"math/rand"
	"sync"
	"sync/atomic"
	"time"
)

func main() {
	// Here, the "state" is a map
	state := make(map[int]int)

	// Mutex will sync access to "state"
	mutex := &sync.Mutex{}

	// Counters for each no. of operations performed
	var readOps, writeOps uint64

	// Here, we perform 100 reads against the state
	for r := 0; r < 100; r++ {
		go func() {
			total := 0
			for {
				// For each read, we pick a key to access,
				key := rand.Intn(5)
				// "Lock" the mutex to ensure exclusive access to the state
				mutex.Lock()
				// Read the value at the chosen key
				total += state[key]
				// "Unlock" the mutex
				mutex.Unlock()
				// Increment the readOps count
				atomic.AddUint64(&readOps, 1)
				// Wait a bit between reads
				time.Sleep(time.Millisecond)
			}
		}()
	}

	// Here, we perform 10 writes against the state
	// using the same patterns we followed for reads.
	for w := 0; w < 10; w++ {
		go func() {
			for {
				key := rand.Intn(5)
				val := rand.Intn(100)
				mutex.Lock()

				state[key] = val
				mutex.Unlock()

				atomic.AddUint64(&writeOps, 1)
				time.Sleep(time.Millisecond)
			}
		}()
	}

	// Let 10 go-routines work on the state & mutex for a second
	time.Sleep(time.Second)

	// Generate the report for read & write operations
	readOpsFinal := atomic.LoadUint64(&readOps)
	fmt.Println("readOps:", readOpsFinal)
	writeOpsFinal := atomic.LoadUint64(&writeOps)
	fmt.Println("writeOps:", writeOpsFinal)

	// a final lock on state, show how it ended up
	mutex.Lock()
	fmt.Println("state:", state)
	mutex.Unlock()
}
