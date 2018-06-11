package main

import "fmt"

func main() {
	records := make([][]string, 0)
	student1 := make([]string, 2)
	student1[0] = "student 1"
	student1[1] = "student 1 value"
	records = append(records, student1)

	student2 := make([]string, 3)
	student2[0] = "student 2"
	student2[1] = "student 2 value"
	student2[2] = "student 2 address"
	records = append(records, student2)

	fmt.Println(records)
}
