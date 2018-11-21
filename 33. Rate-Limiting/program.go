/**
 * Rate Limiting is an important mechanism for controlling resource
 * utilization & maintain quality of service. Go elegantly supports
 * rate limiting with go-routines, channels, and tickers.
 */
package main

import (
	"fmt"
	"time"
)

func main() {
	/**
	 * Firstly we will look at basic rate limiting. Suppose we want
	 * to limit our handling of incoming requests. We will serve
	 * these requests of a channel of the same name.
	 */
	requests := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		requests <- i
	}
	close(requests)

	// This limiter channel will receive a value every 200 ms.
	// This is the regulator in our rate limiting scheme.
	limiter := time.Tick(200 * time.Millisecond)

	// By blocking on receive from the limiter channel before
	// serving each request, we limit ourselves to 1req/200ms.
	for req := range requests {
		<-limiter
		fmt.Println("Request", req, time.Now())
	}

	/**
	 * We may want to allow short bursts of requests in our rate
	 * limiting scheme while preserving the overall rate limit.
	 *
	 * We can accomplish this by buffering our limiter channel.
	 *
	 * This "burstyLimiter" channel will allow bursts of upto 3 events.
	 */
	burstyLimiter := make(chan time.Time, 3)

	// Fill up the channel to represent allowed bursting.
	for i := 0; i < 3; i++ {
		burstyLimiter <- time.Now()
	}

	// Every 200ms we will try to add a new value to "burstyLimiter",
	// upto its limit of 3.
	go func() {
		for t := range time.Tick(200 * time.Millisecond) {
			burstyLimiter <- t
		}
	}()

	/**
	 * Now we simulate 5 more incoming requests. The first 3 of these
	 * will benefit from the burst capability of "burstyLimiter".
	 */
	burstyRequests := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		burstyRequests <- i
	}
	close(burstyRequests)

	for req := range burstyRequests {
		<-burstyLimiter
		fmt.Println("request", req, time.Now())
	}
}
