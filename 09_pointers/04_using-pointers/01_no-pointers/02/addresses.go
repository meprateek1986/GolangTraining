package main

import "fmt"

func main() {
	x := 10
	fmt.Println("x's address is ", &x)
	changeMe(x)
}

func changeMe(x int) {
	fmt.Println("changing value at address ", &x)
	x = 5
}
