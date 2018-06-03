package main

import "fmt"

func main() {
	var country string = "india"
	switch {
	case len(country) == 2:
		fmt.Println("My country name is of length 2")
	case len(country) == 5:
		fmt.Println("My country name is of length 5")
	}
}
