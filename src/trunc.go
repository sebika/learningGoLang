package main

import "fmt"

func main() {
	var x float32

	fmt.Printf("Enter a floating point number: ")
	fmt.Scan(&x)

	y := int32(x)

	fmt.Printf("Truncated number: %d\n", y)
}
