package main

import "fmt"

const metersToYards float64 = 1.09361

func main() {
	var meters float64

	fmt.Print("How many meters you swam : ")
	fmt.Scan(&meters)

	yards := meters * metersToYards

	fmt.Println("you swam ", yards, " yards")
}