package main

import "fmt"

func main() {
	data := []float64{11, 12, 13, 14, 15}
	avg := average(data...)
	fmt.Println(avg)
}

func average(data ...float64) float64 {
	var sum float64
	for _, value := range data {
		sum += value
	}
	return sum / float64(len(data))
}
