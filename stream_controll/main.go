package main

import "fmt"

func main() {
	for i:=1; i<= 10; i++ {
                var str string
		if i%2 ==0 {
			str = "even"
		} else {
			str = "odd"
		}
		fmt.Println(i, str)
	}
}
