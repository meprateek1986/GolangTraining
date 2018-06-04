package main

import "fmt"

func main() {
	fmt.Println(greet("John", "Doe"))
}
func greet(fname string, lname string) (string, string) {
	return "hello " + fname, lname
}
