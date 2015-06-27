package main

import "fmt"

type Person struct {
	Name string
}

func (p *Person) Talk(){
	fmt.Println("Hi, my name is", p.Name)
}

type Android struct {
	Person
	Model string
}

func main() {
	c := Android{Person{"John"},"C11"}
	c.Talk()
}
