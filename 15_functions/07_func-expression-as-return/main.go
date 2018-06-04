package main

import "fmt"

func main() {
	greeter := greeter()
	fmt.Println(greeter())
}
func greeter() func() string {
	return func() string {
		return "hello there"
	}
}
