package main

import (
	"fmt"
)

func main() {
	var num float64
	fmt.Println("\nEnter floating point number:")
	fmt.Scanf("%f", &num)
	var x int = int(num)
	fmt.Printf("\nFloat64 to int gives -> %d", x)
}
