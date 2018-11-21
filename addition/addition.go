package main

import (
	"fmt"
)

func add(x, y int) int {
	return x + y
}

func main() {
	var i, j int = 24, 32
	fmt.Println(add(i, j))
}
