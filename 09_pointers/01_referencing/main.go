package main

import "fmt"

func main() {
	var a int = 42

	fmt.Println(a)
	fmt.Println(&a)

	var b *int = &a

	fmt.Println(b)
}
