package main

import "fmt"

func main() {
	x := 10
	fmt.Println("value of x before changeMe is ", x)
	changeMe(&x)
	fmt.Println("value of x after changeMe is ", x)
}

func changeMe(z *int) {
	*z = 3
}
