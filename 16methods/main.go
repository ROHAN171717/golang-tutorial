package main

import "fmt"

func main() {
	fmt.Println("### Methods in golang ###")

	//no inheritance in golang
	//no super or parent

	raj := User{"Raj", "raj@go.dev", true, 16}
	fmt.Println(raj)
	fmt.Printf("Raj details are: %+v\n", raj)
	fmt.Printf("Name is %v and Email is %v\n", raj.Name, raj.Email)
	raj.GetStatus()
	raj.NewMail()
	fmt.Printf("Name is %v and Email is %v\n", raj.Name, raj.Email)

}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

func (u User) GetStatus() {
	fmt.Println("Is user active: ", u.Status)
}

func (u User) NewMail() {
	u.Email = "test@go.dev"
	fmt.Println("Email of this user is: ", u.Email)
}
