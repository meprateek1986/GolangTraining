package main

import "fmt"

var x int = 10

func main() {
	print()
}

func print() {
	fmt.Println(x)
	x++
	fmt.Println(x)
}
