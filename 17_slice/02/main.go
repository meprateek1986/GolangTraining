package main

import "fmt"

func main() {
	aSlice := make([]int, 0, 5)
	fmt.Println(aSlice)
	fmt.Println(len(aSlice))
	fmt.Println(cap(aSlice))

	for i := 0; i < 80; i++ {
		aSlice = append(aSlice, i)
		fmt.Println(cap(aSlice), " - ", len(aSlice))
	}
}
