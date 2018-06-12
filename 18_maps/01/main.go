package main

import (
	"fmt"
)

func main() {
	m := make(map[string]int)
	m["k1"] = 7
	m["k2"] = 10

	fmt.Println(m)

	delete(m, "k1")

	v, isPresent := m["k1"]

	fmt.Println("Is the value for key k1 present? ", isPresent)
	fmt.Println("value for key k1 is : ", v)

	v1, isK2Present := m["k2"]
	fmt.Println("Is value for key k2 present?", isK2Present)
	fmt.Println("value for key K2 is : ", v1)
}
