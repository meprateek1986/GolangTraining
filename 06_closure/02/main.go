package main

import "fmt"

func print() {
	var x int = 10
	{
		fmt.Println(x)
		x++
		fmt.Println(x)
	}
}

func main() {
	print()
}
