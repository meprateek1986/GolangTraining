package main

import "fmt"

func main() {
	greeting := func() { //assigning a function to a variable is called func expression
		fmt.Println("hello there")
	}
	greeting()
}
