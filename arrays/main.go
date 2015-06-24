package main

import "fmt"
func main() {
	var x [5]float64
	x[0] = 54
	x[1] = 77
	x[2] = 81
	x[3] = 98
	x[4] = 51

	var total float64 = 0
	for _, value := range x {
		total += value
	}
	fmt.Println(total/float64(len(x)))
}
