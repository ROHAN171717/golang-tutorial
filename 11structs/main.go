package main

import "fmt"

func main() {
	fmt.Println("### Structs in golang ###")

	//no inheritance in golang
	//no super or parent

	raj := User{"Raj", "raj@go.dev", true, 16}
	fmt.Println(raj)
	fmt.Printf("Raj details are: %+v\n", raj)
	fmt.Printf("Name is %v and Email is %v", raj.Name, raj.Email)

}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}
