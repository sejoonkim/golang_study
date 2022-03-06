package main

import "fmt"

func main() {
	fmt.Println("Structs in golang")

	// no inheritance in golang: No super or parent

	sejoon := User{"Sejoon", "sejoon@dif.com", true, 16, 1}
	fmt.Println(sejoon)
	fmt.Printf("sejoon details are: %+v\n", sejoon)
	fmt.Printf("Name is %v and email is %v\n", sejoon.Name, sejoon.Email)
	sejoon.GetStatus()
	sejoon.NewMail()
	fmt.Printf("Name is %v and email is %v\n", sejoon.Name, sejoon.Email)
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
	oneAge int
}

func (u User) GetStatus() {
	fmt.Println("Is user active:", u.Status)

}

func (u User) NewMail() {
	u.Email = "test@go.dev"
	fmt.Println("Email of this user is:", u.Email)
}
