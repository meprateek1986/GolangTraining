package main

import (
	"fmt"
)

func main() {
	m := map[int]string{
		0: "hello",
		1: "world",
	}

	for k, val := range m {
		fmt.Println("key is ", k, " and value is ", val)
	}
}
