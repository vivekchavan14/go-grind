package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("userInput")
	welcome := "Welcome to"
	fmt.Println(welcome)
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter an input: ")
	input, _ := reader.ReadString('\n')
	fmt.Println("UserInput: ", input)

}
