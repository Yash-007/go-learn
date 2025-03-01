package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

func TestGet() {
	const url = "http://localhost:8080"
	response, err := http.Get(url)

	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	fmt.Println("status code: ", response.StatusCode)
	fmt.Println("status: ", response.Status)
	fmt.Println("content length: ", response.ContentLength)

	content, _ := io.ReadAll(response.Body)
	fmt.Println(string(content))
}

func TestPost() {
	const url = "http://localhost:8080"
	payload := strings.NewReader(`
	  {
	   "name": "Yash",
	   "hobby": "Reading"
	  }
	`)
	response, err := http.Post(url, "application/json", payload)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()
	content, _ := io.ReadAll(response.Body)
	fmt.Println(string(content))
}
func TestPostForm() {
	const myurl = "http://localhost:8080"

	data := url.Values{}
	data.Add("name", "Yash")
	data.Add("hobby", "Reading")

	response, err := http.PostForm(myurl, data)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()
	content, _ := io.ReadAll(response.Body)
	fmt.Println(string(content))
}

func main() {
	// TestGet()
	TestPostForm()
}
