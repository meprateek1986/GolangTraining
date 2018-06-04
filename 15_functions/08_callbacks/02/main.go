package main

import "fmt"

func main() {
	var a []int = []int{1, 2, 3, 10}
	for _, value := range a {
		print(value, func(v int) {
			fmt.Println(v)
		})
	}
}

func print(val int, callback func(v int)) {
	callback(val)
}
