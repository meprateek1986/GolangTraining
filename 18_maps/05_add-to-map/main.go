package main

import (
	"fmt"
)

func main() {
	var m = make(map[string]string)
	m["name"] = "John"
	m["passion"] = "wrestling"

	fmt.Println(m)
}
