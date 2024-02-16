package main

import "fmt"

func main() {
	fmt.Println("### Loops in golang ###")

	days := []string{"Sun", "Mon", "Tue", "Wed", "Thu"}

	fmt.Println(days)

	for d := 0; d < len(days); d++ {
		fmt.Println(days[d])
	}

	fmt.Println()

	for i := range days {
		fmt.Println(days[i])
	}

	fmt.Println()

	for index, day := range days {
		fmt.Printf("Index is %v and Value is %v\n", index, day)
	}

	val := 1

	for val < 10 {
		if val == 9 {
			goto lco
		}
		if val == 5 {
			val++
			continue
		}
		fmt.Println("value is: ", val)
		val++
	}

lco:
	fmt.Println("Jumping using goto...")
}
