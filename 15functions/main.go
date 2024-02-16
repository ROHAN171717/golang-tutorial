package main

import "fmt"

func main() {
	fmt.Println("### Function in golang ###")

	greeter()

	result := adder(3, 5)
	fmt.Println("Result: ", result)

	proRes, msg := proAdder(2, 5, 8, 7, 3)
	fmt.Println("Pro result is: ", proRes)
	fmt.Println("Pro message is: ", msg)

}

func greeter() {
	fmt.Println("Hello...")
}

func adder(x int, y int) int {
	return x + y
}

func proAdder(values ...int) (int, string) {
	total := 0

	for _, val := range values {
		total += val
	}

	return total, "How are you?"
}
