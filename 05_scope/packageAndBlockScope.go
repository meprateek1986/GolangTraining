package main

import (
	"fmt"

	"github.com/GolangTraining/05_scope/differentpackage"
)

func main() {
	var blockLevelScope1 int = 4
	fmt.Println("block level variable value is - ", blockLevelScope1)
	fmt.Println("package level variable value is - ", differentpackagedeclaration.PackageLevelScope1)
	differentpackagedeclaration.PrintFromPackage()
}

func block() {
	fmt.Println("not able to access block level variable outside block scope")
}
