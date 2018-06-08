package main

import "fmt"

func main() {
	var x [8]int
	fmt.Println(x)
	x[7] = 3
	fmt.Println(x)
}
