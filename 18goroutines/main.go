package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
)

var wg sync.WaitGroup
var mut sync.Mutex

var signals []string

func main() {
	// go greeters("Hello")
	// greeters("world")

	var urls = []string{
		"https://whatsapp.com",
		"https://linkedin.com",
		"https://facebook.com",
		"https://whatsapp.com",
	}

	for _, res := range urls {
		go getRes(res)
		wg.Add(1)
	}
	wg.Wait()
}

// func greeters(s string) {
// 	for i := 0; i < 5; i++ {
// 		time.Sleep(3 * time.Millisecond)
// 		fmt.Println(s)
// 	}
// }

func getRes(url string) {
	defer wg.Done()
	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	} else {
		mut.Lock()
		signals = append(signals, url)
		mut.Unlock()

		fmt.Printf("status code received: %d for url: %s\n\n", res.StatusCode, url)
	}
}
