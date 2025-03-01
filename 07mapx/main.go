package main

import "fmt"

func main() {
	// mp := map[string]string{}
	mp := make(map[string]string)
	mp["shivam"] = "Yash"
	mp["raj"] = "Raj"

	for name, actName := range mp {
		fmt.Printf("The name is %v and actual name is %v\n", name, actName)
	}
}
