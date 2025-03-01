package main

import "fmt"

type Person struct {
	Name string
	Age  int
	Mob  int
}

func main() {
	yash := Person{"Yash", 21, 134343}
	fmt.Printf("Yash is %+v %+v %+v", yash.Name, yash.Age, yash.Mob)
}
