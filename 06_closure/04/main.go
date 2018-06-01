package main

import "fmt"

func print() func() int {
	var x int = 10
	return func() int {
		x++
		return x
	}
}

func main() {
	increment := print()
	fmt.Println(increment())
	fmt.Println(increment())
}
