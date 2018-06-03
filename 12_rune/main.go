package main

import "fmt"

func main() {
	a := 'a'               //a rune is a character and the value gets stored as int32
	fmt.Println(a)         //this will print 97 as character 'a' is represented by 97 in UTF-8
	fmt.Printf("%T \n", a) //type of a rune is int32
}
