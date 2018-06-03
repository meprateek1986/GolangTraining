package main

import "fmt"

func main() {
	a := 33
	var b *int = &a

	*b = 44 //changing value at the address stored inside b. This will change the value of a

	fmt.Println(a) //printing out changed value of a
}
