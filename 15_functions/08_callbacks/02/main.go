package main

import "fmt"

func main() {
	var a []int = []int{1, 2, 3, 10}
	xs := print(a, func(v int) bool {
		return v > 1
	})
	fmt.Println(xs)
}

func print(values []int, callback func(v int) bool) []int {
	xs := []int{}
	for _, val := range values {
		if callback(val) {
			xs = append(xs, val)
		}
	}
	return xs
}
