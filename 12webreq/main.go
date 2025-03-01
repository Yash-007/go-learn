package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	const url = "https://google.com"
	res, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
	}

	defer res.Body.Close()

	dataBytes, _ := io.ReadAll(res.Body)
	fmt.Print(string(dataBytes))

}
