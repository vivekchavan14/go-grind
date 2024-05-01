package main

import "fmt"

func main() {
	fmt.Println("Learning pointers!")
	var ptr *int
	fmt.Println("Default value of pointer is : ", ptr)
	myNum := 23
	var potr = &myNum
	fmt.Println("Value of pointer potr is :", potr)
	fmt.Println("Value of pointer potr is :", *potr)
}
