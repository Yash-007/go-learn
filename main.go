package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello this is yash.")
	var a string
	fmt.Println("a is: ", a)
	a = "test"
	fmt.Println("a is: ", a)
	fmt.Printf("a is of type %T \n", a)

	fmt.Println("Enter the rating for the pizza: ")
	// reader := bufio.NewReader(os.Stdin)

	// input, _ := reader.ReadString('\n')
	input := "343"

	fmt.Println("Thanks for rating: ", input)
	fmt.Printf("Type: %T", input)
}
