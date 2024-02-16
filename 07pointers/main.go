package main

import "fmt"

func main() {
	fmt.Println("Welcome to time study of pointers...")

	// var ptr *int 
	// fmt.Println("Value of pointer is: ", ptr)

	x := 23;

	var ptr = &x;

	fmt.Println("Value of pointer is: ", ptr)
	fmt.Println("Value of pointer is: ", *ptr)

	*ptr = *ptr + 2;
	fmt.Println("New Value is: ", x)
}