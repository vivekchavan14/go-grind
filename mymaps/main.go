package main

import "fmt"

func main() {
	fmt.Println("Maps in Go!")
	languages := make(map[string]string)
	languages["JS"] = "Javascript"
	languages["PHP"] = "PHP"
	fmt.Println("List of all languages : ", languages)
	//delete language
	delete(languages, "PHP")
	fmt.Println("List of all languages : ", languages)

	//looping through maps
	for key, value := range languages {
		fmt.Printf("For key %v, value %v \n", key, value)
	}
}
