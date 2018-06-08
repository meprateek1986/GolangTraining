package main

import "fmt"

func main() {
	aSlice := make([]int, 0, 5)
	aSlice = append(aSlice, 1)
	aSlice = append(aSlice, 2)
	aSlice = append(aSlice, 3)
	aSlice = append(aSlice, 4)
	aSlice = append(aSlice, 5)

	fmt.Println(aSlice)

	//deleting element 3 from aSlice by appending aSlice from 0 to 3(exclusive) and then from 4 to the end. So ultimately deleting element at position 3
	aSlice = append(aSlice[0:3], aSlice[4:]...)

	fmt.Println(aSlice)
}
