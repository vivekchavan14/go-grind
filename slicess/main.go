package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Slices in Go")
	var numList = []int{}
	numList = append(numList, 0, 1, 2, 3, 4)
	fmt.Println("NUM LIST IS : ", numList)
	numList = append(numList[1:])
	fmt.Println("Now the num list in slices is :", numList)
	numList = append(numList[1:4])
	fmt.Println("Now the num list in slices is :", numList)

	keyList := make([]int, 2)
	keyList[0] = 88
	keyList[1] = 34
	fmt.Println("Key list is : ", keyList)
	keyList = append(keyList, 23, 44, 55, 67)
	fmt.Println("Updated key list is : ", keyList)

	//sort package sorts integers, float values, strings etc
	sort.Ints(keyList)
	fmt.Println(keyList)
}
