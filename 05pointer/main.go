package main

import "fmt"

func main() {
	// fmt.Println("welcome pointers.")
	var FL [4]string

	FL[0] = "apple"
	fmt.Println(len(FL))
	fmt.Println(FL)
}
