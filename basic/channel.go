package basic

import "fmt"

// ChannelSelect ...
func ChannelSelect(c, quit chan int) {
	x, y := 0, 1
	for { // as while loop
		select {
		// evaluate c <- x to run
		// also see, try send or try block operations
		// https://go101.org/article/channel.html
		// https://go101.org/article/channel-use-cases.html
		case c <- x:
			x, y = y, x+y
		case <-quit:
			fmt.Println("quit channel get piped")
			return
		}
	}
}

// ChannelRange ...
func ChannelRange(ch chan int) {
	for i := range ch {
		fmt.Println(i)
	}
}

// ChannelFib ...
func ChannelFib(n int, ch chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		ch <- x
		x, y = y, x+y
	}
	close(ch)
}
