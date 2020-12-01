package basic

import (
	"fmt"
	"time"
)

// RunSelectWithTimeout ...
func RunSelectWithTimeout() {
	tick := time.Tick(100 * time.Millisecond)
	boom := time.After(500 * time.Millisecond)
	count := 0
	for {
		select {
		case <-tick:
			fmt.Println("tick ", count, " times")
			count = 0
		case <-boom:
			fmt.Println("BOOM!")
			return
		default:
			// this default loop will runs inside for loop
			// if no cases matched, fall into this
			count = count + 1
		}
	}
}
