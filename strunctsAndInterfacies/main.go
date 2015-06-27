package main

import ("fmt"; "math")

type Circle struct {
	x, y, r float64
}

func circleArea(c Circle) float64 {
	return math.Pi * c.r*c.r
}

func main() {
	var one Circle
	two := new(Circle)
	three := Circle{x: 0, y: 0, r: 5}
	four := Circle{0, 0, 5}
	fmt.Println(circleArea(one))
	fmt.Println(circleArea(*two))
	fmt.Println(circleArea(three))
	fmt.Println(circleArea(four))
}
