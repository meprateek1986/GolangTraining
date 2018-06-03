package main

import "fmt"

func main() {
	switch "Fun" {
	case "Joy":
		fmt.Println("I am having Joy")
	case "Fun":
		fmt.Println("I am having Fun")
	default:
		fmt.Println("I am in default mood :D")
	}
}
