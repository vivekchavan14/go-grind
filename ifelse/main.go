package main

import "fmt"

func main() {
	fmt.Println("Ifelse in Go")
	count := 20
	var str string
	if count > 10 {
		str = "number is greater than 10"
	}
	fmt.Println(str)
}
