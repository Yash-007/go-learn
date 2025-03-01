package main

import (
	"fmt"
	"time"
)

func main() {
	tm := time.Now()
	fmt.Println(time.Now())
	fmt.Println(tm.Format("01-02-2006 Monday 15:04:05"))

	date := time.Date(2020, time.August, 30, 2, 34, 23, 23, time.Local)
	fmt.Println(date.Format("01-02-2006, Monday, 15:04:05"))

}
