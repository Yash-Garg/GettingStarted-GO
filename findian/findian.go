package main

import (
	"fmt"
	"strings"
)

func main() {
	var str string
	fmt.Print("\nEnter a string : ")
	fmt.Scan(&str)
	if (strings.HasPrefix(str, "i") || strings.HasPrefix(str, "I")) && (strings.HasSuffix(str, "n") || strings.HasSuffix(str, "N")) {
		if strings.Contains(str, "a") || strings.Contains(str, "A") {
			fmt.Print("\nFound!")
		}
	} else {
		fmt.Print("\nNot Found!")
	}
}
