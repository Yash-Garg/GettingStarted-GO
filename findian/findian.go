package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewScanner(os.Stdin)
	fmt.Println("\nEnter a string : ")

	reader.Scan()

	input := strings.ToLower(reader.Text())

	if (strings.HasPrefix(input, "i")) && (strings.HasSuffix(input, "n")) && (strings.Contains(input, "a")) {
		fmt.Println("\nFound!")
	} else {
		fmt.Println("\nNot Found!")
	}
}
