package main

import (
	"fmt"
)

func main() {
	m := map[string]string{}
	m["k1"] = "hello"
	fmt.Println(m)
	delete(m, "k1")
	fmt.Println(m)
}
