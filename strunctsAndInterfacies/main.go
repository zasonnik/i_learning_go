package main

import "fmt"

type Circle struct {
	x, y, r float64
}

func main() {
	var one Circle
	two := new(Circle)
	three := Circle{x: 0, y: 0, r: 5}
	four := Circle{0, 0, 5}
	fmt.Println(one.x, one.y, one.r)
	fmt.Println(two.x, two.y, two.r)
	fmt.Println(three.x, three.y, three.r)
	fmt.Println(four.x, four.y, four.r)
}
