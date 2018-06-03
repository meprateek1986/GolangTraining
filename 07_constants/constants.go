package main

import "fmt"

const constWithDeclaration string = "string constant with declaration"
const constWitoutDeclaration = "string constant without declaration"
const (
	a = iota //0
	b = iota //1
)

func main() {
	fmt.Println(constWithDeclaration)
	fmt.Println(constWitoutDeclaration)
	fmt.Println(a)
	fmt.Println(b)

	const x = 42            //untyped constant
	const xint int = 33     //typed constant
	const y float64 = 44.23 //typed constant

	z := x + xint //this will work because type of x will be defined at compile time and will be given type of other typed constant
	z1 := x + y   //this will work because type of x will be defined at compile time and will be given type of other typed constant
	//z2 := xint + y	//this can not be done. as both xint and y are typed and two different types can not be mixed in go

	fmt.Println(z)
	fmt.Println(z1)
}
