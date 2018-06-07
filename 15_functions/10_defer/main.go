package main

import "fmt"

func main() {
	//defer keyword will make the execution of the method right before the end of the method it is enclosed with(main() in this case)
	//this can be used when user wants to write the code close to where it makes sense but still wants to execute it towards the end
	//for eg. if a file is opened it needs to be closed before the end of the method but if we write that code to close the file near
	//the line where we are opening it then it will be easily visible and file close will happen right before the method exits
	defer world()
	hello()
}

func hello() {
	fmt.Println("hello")
}

func world() {
	fmt.Println("world")
}
