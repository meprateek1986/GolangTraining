package main

import "fmt"

func main() {
	switch "truck" {
	case "car", "bike":
		fmt.Println("I can drive truck or bike")
	case "plane", "truck":
		fmt.Println("I can fly a plane or drive a truck")
	}
}
