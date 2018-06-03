package main

import "fmt"

func main() {
	switch "toyota" {
	case "maruti":
		fmt.Println("I am driving maruti")
	case "honda":
		fmt.Println("I am driving honda")
		fallthrough
	case "toyota":
		fmt.Println("I am driving toyota")
		fallthrough
	case "tata":
		fmt.Println("I am driving tata")
	}
}
