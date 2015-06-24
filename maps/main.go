package main

import "fmt"

func main(){
	x := make(map[string]int)
	x["key"] = 10
	x["key1"] = 20
	delete(x,"key")
	v,ok := x["key1"]
	if ok {
		fmt.Println(v)
	} else {
		fmt.Println("nope")
	}
}
