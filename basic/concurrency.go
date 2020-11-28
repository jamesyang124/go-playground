package basic

import (
	"fmt"
	"sync"
)

// GoroutinePrint ...
func GoroutinePrint(wg *sync.WaitGroup, s string) {
	defer wg.Done()
	fmt.Println(fmt.Sprintf("goroutine run this - %s", s))
}

// GoroutinePrintf ...
func GoroutinePrintf(wg *sync.WaitGroup, s string) string {
	defer wg.Done()
	return fmt.Sprintf("goroutine run this - %s", s)
}

// ChannelPass ...
func ChannelPass(val int, ch chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	ch <- val + 5
}
