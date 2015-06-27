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

type Shape interface {
	area() float64
}

func totalArea(shapes ...Shape) float64 {
	var area float64
	for _, s := range shapes {
		area += s.area()
	}
	return area
}

func main() {
	c := Circle{0, 0, 5}
	r := Rectangle{0,0,3,4}
	fmt.Println(totalArea(&c,&r))
}
