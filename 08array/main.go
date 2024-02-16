package main

import "fmt"

func main() {
	fmt.Println("### Array in golang ###")

	var arr [4]string

	arr[0] = "a"
	arr[1] = "b"
	arr[3] = "c"

	fmt.Println("Array: ", arr)
	fmt.Println("Array Length: ", len(arr))

	var arr2 = []string{"a", "b", "c"}
	fmt.Println("Array Length: ", len(arr2))
}
