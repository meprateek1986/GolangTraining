package main

import (
	"fmt"
)

func main() {
	var m map[string]string
	fmt.Println("is map nil? ", m == nil)

	//adding value to nil map will throw an error at runtime so never do that
	m["hello"] = "world"
}
