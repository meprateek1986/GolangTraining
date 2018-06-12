package main

import (
	"fmt"
)

func main() {
	var m = make(map[string]string)
	m["name"] = "John"
	fmt.Println(m)
	m["name"] = "Cena"
	fmt.Println(m)
}
