package main

import (
	"fmt"
	"bufio"
)

func main() {

    var age int
    fmt.Println("What is your age?")
    _, err: fmt.Scan(&age)												    // Reading an integer

    reader := bufio.newReader(os.Stdin)
    var name string
    fmt.Println("What is your name?")
    name, _ := reader.readString("\n")										// Reading a string

    fmt.Println("Your name is ", name, " and your age is ", age)
}
