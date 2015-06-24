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
	for i:=0; i<len(x); i++ {
		total +=x[i]
	}
	fmt.Println(total/float64(len(x)))
}
