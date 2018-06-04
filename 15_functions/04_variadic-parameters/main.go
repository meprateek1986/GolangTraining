package main

import "fmt"

func main() {
	avg := average(11, 12, 13, 14, 15)
	fmt.Println(avg)
}
func average(vals ...float64) float64 { //... are called variadic. in java its called varargs
	var sum float64
	for _, value := range vals {
		sum += value
	}
	return sum / float64(len(vals))
}
