package main

import "fmt"

func main() {
	loop([]int{1, 2, 3, 4}, func(j int) {
		fmt.Println(j)
	})
}

func loop(data []int, callback func(n int)) {
	for _, value := range data {
		callback(value)
	}
}
