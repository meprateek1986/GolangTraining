package main

import (
	"fmt"
)

func main() {
	var m = make(map[int]string, 1)
	m[0] = "hello"
	m[1] = "world"

	if v, ok := m[1]; !ok {
		fmt.Println("value does not exists for key 1")
	} else {
		fmt.Println("value for key 1 is ", v)
	}
}
