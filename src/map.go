package main

import "fmt"

type Person struct {
	name  string
	addr  string
	phone string
}

func main() {
	myMap := make(map[string]int)
	myMap["Ionel"] = 1
	myMap["Gigel"] = 2
	myMap["Purcel"] = 3

	p1 := new(Person)
	p1.name = "Georgel"
	p1.addr = "Str Str"
	p1.phone = "34-9023-948"
	fmt.Println(p1)

	for key, value := range myMap {
		fmt.Printf("key: %s value: %d\n", key, value)
	}

}
