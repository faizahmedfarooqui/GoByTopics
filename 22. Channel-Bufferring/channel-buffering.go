/**
 * By default, channels are unbuffered.
 *
 * Meaning that they will only accept sends (chan <-) if there is
 * corresponding receive (<- chan) ready to receive the sent value.
 *
 * Buffered channels accept a limited number of values without a
 * corresponding receiver for those values.
 */
package main

import (
	"fmt"
)

func main() {
	// Create channel syntax with buffered limit value
	messages := make(chan string, 2)

	messages <- "buffered"
	messages <- "channel"

	fmt.Println(<-messages)
	fmt.Println(<-messages)
}
