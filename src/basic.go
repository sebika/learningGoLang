package main

import "fmt"

func main() {
	vect := [...]int{1, 2, 3, 4, 5}

	for _, v := range vect {
		v += 1
	}

	for _, v := range vect {
		fmt.Printf("%d ", v)
	}
}
