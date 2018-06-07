package main

import "fmt"

func main() {
	var m []string = make([]string, 1, 10)
	m[0] = "hello"
	fmt.Println(m)
	changeMe(m)
	fmt.Println(m)
}

func changeMe(z []string) {
	z[0] = "world"
}
