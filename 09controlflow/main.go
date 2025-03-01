package main

import "fmt"

func main() {
	// var a = (rand.Intn(6) + 1)
	// switch a {
	// case 1:
	// 	fmt.Println("it is one")
	// case 2:
	// 	fmt.Println("it is two")
	// case 3:
	// 	fmt.Println("it is three")
	// 	fallthrough
	// case 4:
	// 	fmt.Println("it is four")
	// case 5:
	// 	fmt.Println("it is five")
	// case 6:
	// 	fmt.Println("it is six")
	// default:
	// 	fmt.Println("it is nothing")
	// }

	flist := []int{1, 2, 3, 4, 5}
	// for i, val := range flist {
	// 	fmt.Printf("index is %v and value is %v\n", i, val)
	// }
	for val := 0; val < len(flist); val++ {
		fmt.Printf("The value is %v\n", val)
	}

}
