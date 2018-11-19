/**
 * Closing a channel indicates that no more values will be sent on it.
 * This can be useful to communicate completion to the channel's receivers.
 *
 * Here, we have an example where we will use a "jobs" channel to communicate
 * work to be done from the "main()" go-routine to a "worker" go-routine.
 * When we have no more jobs for the worker we will close the "jobs" channel.
 */
package main

import (
	"fmt"
)

func main() {
	jobs := make(chan int, 5)
	done := make(chan bool)

	// Worker GoRoutine
	go func() {
		for {
			// Here, in this special 2-value form of receive, the
			// "more" value will be "false" if "jobs" has been "closed"
			// and all values in the channel have already been received.
			// It is used to notify on "done" when we've worked on all "jobs".
			j, more := <-jobs
			if more {
				fmt.Println("received job", j)
			} else {
				fmt.Println("received all jobs")
				done <- true
				return
			}
		}
	}()

	// From here, we will send 2 jobs to the worker
	// over the "jobs" channel then closes it.
	for j := 1; j <= 3; j++ {
		jobs <- j
		fmt.Println("sent job", j)
	}
	close(jobs)
	fmt.Println("sent all jobs")

	// We await the worker using the
	// channel-sync approach
	<-done
}
