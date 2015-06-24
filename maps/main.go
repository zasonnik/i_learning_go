package main

import "fmt"

func main(){
	x := make(map[string]int)
	x["key"] = 10
	x["key1"] = 20
	delete(x,"key")
	fmt.Println(x)
}
