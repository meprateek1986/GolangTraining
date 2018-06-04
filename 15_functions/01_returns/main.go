package main

import "fmt"

func main() {
	fmt.Println(greet("hello", "dev"))
}

func greet(greeting, name string) string {
	return fmt.Sprint(greeting, name)
}
