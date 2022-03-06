package main

import "fmt"

func main() {
	fmt.Println("Structs in golang")

	// no inheritance in golang: No super or parent

	sejoon := User{"Sejoon", "sejoon@dif.com", true, 16}
	fmt.Println(sejoon)
	fmt.Printf("sejoon details are: %+v\n", sejoon)
	fmt.Printf("Name is %v and email is %v\n", sejoon.Name, sejoon.Email)

}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}
