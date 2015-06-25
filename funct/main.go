
package main

import "fmt"

func first() {
	fmt.Println("1st")
}

func second() {
	fmt.Println("2st")
}

func main() {
	defer second()
	first()
	fmt.Println("middle")
}
