package main

import (
	"fmt"
	"sync"
)

// https://chatgpt.com/share/67db2ab2-eab8-8007-b5c0-4346a86ba6a5 (at the end)

// channels are needed for:
// 1. communication between the channels
// 2. synchronization for execution (can also use time.sleep and wait groups)
func main() {
	fmt.Println("Channels in Go")

	myCh := make(chan int, 1) // it means we can take one extra value in addition to normal one so in this case if we send two values in channel, atleast one needs to be recieved
	wg := &sync.WaitGroup{}

	// myCh <- 5
	// val := <-myCh
	// fmt.Println(val)

	wg.Add(2)

	go func(ch <-chan int, wg *sync.WaitGroup) {
		val, isChannelOpen := <-ch

		fmt.Println(isChannelOpen)
		fmt.Println(val)
		fmt.Println("working fine")

		wg.Done()
	}(myCh, wg)

	go func(ch chan<- int, wg *sync.WaitGroup) {
		ch <- 0
		ch <- 2
		close(ch)
		wg.Done()
	}(myCh, wg)

	wg.Wait()
}
