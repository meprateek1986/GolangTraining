package main

import "fmt"

func main() {
	max := findMax(1, 2, 3, 4, 5)
	fmt.Println(max)
}

func findMax(a ...int) int {
	var max int
	for _, value := range a {
		if value > max {
			max = value
		}
	}
	return max
}
