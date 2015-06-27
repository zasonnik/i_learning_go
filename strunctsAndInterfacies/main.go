package main

import ("fmt"; "math")

type Circle struct {
	x, y, r float64
}

func (c *Circle) area() float64 {
	return math.Pi * c.r*c.r
}

type Rectangle struct {
	x1, y1, x2, y2 float64
}

func (r *Rectangle) area() float64 {
	l := math.Abs(r.x1-r.x2)
	w := math.Abs(r.y1-r.y2)
	return l * w
}

func main() {
	c := Circle{0, 0, 5}
	r := Rectangle{0,0,3,4}
	fmt.Println(c.area())
	fmt.Println(r.area())
}
