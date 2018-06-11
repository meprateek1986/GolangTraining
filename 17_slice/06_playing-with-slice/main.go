package main

import "fmt"

func main() {
	aSlice := []string{}               //empty string slice
	var newSlice = make([]string, 5)   //making a string slice with length 5. This will by default choose capacity 5.
	var nSlice = make([]string, 5, 15) //making a slice of length 5 and capacity 15

	fmt.Println(aSlice == nil)
	fmt.Println(cap(newSlice))
	fmt.Println(cap(nSlice))
}
