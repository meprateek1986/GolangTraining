package main

import "fmt"

func main() {
	aSlice := []int{1, 2, 3, 4, 5}
	fmt.Println(aSlice[2:4])
	fmt.Println(aSlice[2:])
	fmt.Println(aSlice[:4])
	fmt.Println(aSlice[:])
}
