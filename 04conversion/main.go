package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Welcome to our pizza app")

	reader := bufio.NewReader(os.Stdin)
	num, _ := reader.ReadString('\n')
	newNum, _ := strconv.ParseInt(strings.TrimSpace(num), 10, 64)
	fmt.Println(newNum + 1)
}
