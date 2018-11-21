/**
 * Worker is implemented using go-routines & channels.
 */
package main

import (
	"fmt"
	"time"
)

/**
 * This function will run several concurrent instances.
 *
 * These workers will receive work on the "jobs" channel
 * and send the corresponding results on the "results".
 *
 * We will sleep a second per job to simulate an
 * expensive task.
 */
func worker(id int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		fmt.Println("Worker", id, "started job", j)
		time.Sleep(time.Second)

		fmt.Println("Worker", id, "finished job", j)
		results <- j * 2
	}
}

func main() {
	// In order to use our pool of workers, we need to send
	// them work & collect their results.
	jobs := make(chan int, 100)
	results := make(chan int, 100)

	// This starts-up 3 workers, initially blocked because
	// there are no jobs yet.
	for w := 1; w <= 3; w++ {
		go worker(w, jobs, results)
	}

	// Here, we send 5 jobs & then close "jobs" channel
	for j := 1; j <= 5; j++ {
		jobs <- j
	}
	close(jobs)

	// Finally we collect all the results of the work
	for a := 1; a <= 5; a++ {
		<-results
	}
}
