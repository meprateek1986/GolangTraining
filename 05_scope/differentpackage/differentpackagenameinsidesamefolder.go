package differentpackagedeclaration

import "fmt"

func PrintFromPackage() {
	fmt.Println(notExportedPackageLevelScope)
}
