package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("Channels in Go")

	myCh := make(chan int, 2)
	wg := &sync.WaitGroup{}

	// myCh <- 5
	// val := <-myCh
	// fmt.Println(val)

	wg.Add(2)

	go func(ch <-chan int, wg *sync.WaitGroup) {
		val, isChannelOpen := <-ch

		fmt.Println(isChannelOpen)
		fmt.Println(val)

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
