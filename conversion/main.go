package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Welcome to pizza app")
	fmt.Println("Rate our pizza between 1 && 5")
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input) // Trim leading and trailing whitespace, including newline characters
	fmt.Println("Thanks for rating us ", input)

	newRating, err := strconv.ParseFloat(input, 64)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Rating is updated to ", newRating+1)
	}
}
