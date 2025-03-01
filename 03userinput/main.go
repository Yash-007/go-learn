package main

import (
	"fmt"
	"log"
)

func main() {
	var i int
	fmt.Println("enter i: ")
	_, err := fmt.Scan(&i)
	if err != nil {
		log.Print("error", err)
	}
	fmt.Printf("Type of i is %T \n", i)
	fmt.Println("i is ", i+1)

	fmt.Println("This is main.go inside user input")
}
