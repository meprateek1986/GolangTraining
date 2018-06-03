package main

import "fmt"

func main() {
	x := 10
	changeMe(x)
	fmt.Println(x)
}

func changeMe(x int) {
	x = 0
}
