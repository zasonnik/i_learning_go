package main

import "fmt"

func main() {
	fmt.Print("Введите число ");
	var input float32
	fmt.Scanf("%f", &input)

	var out float64 = float64(input * 2)

	fmt.Println(out)
}
