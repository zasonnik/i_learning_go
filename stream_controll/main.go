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
		var str2 string
                switch i {
                    case 0: str2 = "Zero"
                    case 1: str2 = "One"
                    case 2: str2 = "Two"
                    default: str2 = "Too mutch"
                }
		fmt.Println(i, str, str2)
	}
}
