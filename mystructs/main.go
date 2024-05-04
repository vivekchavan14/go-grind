package main

import "fmt"

func main() {
	fmt.Println("Structs in GO!")
	goEngineer := User{"vivek", "vivek@go.dev", true, 21}
	fmt.Println(goEngineer)
	fmt.Printf("Details of Go engineer : %+v \n", goEngineer)
	fmt.Printf("Name : %v && email : %v \n", goEngineer.Name, goEngineer.Email)
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}
