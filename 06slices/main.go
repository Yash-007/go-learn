package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("abcd")
	fruits := []int{2, 3}
	fmt.Println(fruits)

	sliceList := make([]int, 4)
	sliceList[0] = 1
	sliceList[1] = 2

	fmt.Println(sort.IntsAreSorted(sliceList))
	fmt.Println(sort.IsSorted(sort.IntSlice{2, 3, 4, 6}))
	// sliceList = append(sliceList, 2, 3)
	fmt.Println(sliceList)

	vegies := []int{2, 3, 5, 7, 6, 2}
	vegies = append(vegies[:2], vegies[3:]...)
	fmt.Println(vegies)

}
