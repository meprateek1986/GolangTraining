package main

import "fmt"

func print() {
	var x int = 10

	increment := func() int {
		x++
		return x
	}

	fmt.Println(increment())
	fmt.Println(increment())
}

func main() {
	print()
}
