package main

import "fmt"

func main() {
	d, b := half(1)
	fmt.Println(d, ", ", b)

	d, b = half(2)
	fmt.Println(d, ", ", b)
}

func half(i int) (int, bool) {
	var d int = i / 2
	var b bool = (i%2 == 0)
	return d, b
}
