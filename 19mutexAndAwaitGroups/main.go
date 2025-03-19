package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("Race Condition.")

	wg := &sync.WaitGroup{}
	mut := &sync.Mutex{}

	var score = []int{0}

	wg.Add(2)

	go func(w *sync.WaitGroup, m *sync.Mutex) {
		fmt.Println("Two R")
		// m.Lock()
		m.Unlock()
		score = append(score, 2)
		wg.Done()
	}(wg, mut)

	go func(w *sync.WaitGroup, m *sync.Mutex) {
		fmt.Println("One R")
		m.Lock()
		score = append(score, 1)
		// m.Unlock()
		wg.Done()
	}(wg, mut)

	// go func(w *sync.WaitGroup, m *sync.Mutex) {
	// 	fmt.Println("Three R")
	// 	m.Lock()
	// 	score = append(score, 3)
	// 	m.Unlock()
	// 	wg.Done()
	// }(wg, mut)

	wg.Wait()

	fmt.Println(score)
}
