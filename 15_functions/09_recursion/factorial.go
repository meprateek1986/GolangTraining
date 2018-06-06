package main

import "fmt"

func main() {
	fact := factorial(4)
	fmt.Println(fact)
}

func factorial(i int) int {
	if i == 0 {
		return 1
	}

	return i * factorial(i-1)
}
