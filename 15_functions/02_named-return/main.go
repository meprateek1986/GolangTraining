package main

import "fmt"

func main() {
	s := greet("hello", "dev")
	fmt.Println(s)
}

func greet(greeting string, name string) (s string) {
	s = fmt.Sprint(greeting, " ", name)
	return
}
