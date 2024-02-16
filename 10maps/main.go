package main

import "fmt"

func main() {
	fmt.Println("### Maps in golang ###")

	m := make(map[string]string)

	m["1"] = "aaa"
	m["2"] = "bbb"
	m["3"] = "ccc"

	fmt.Println("key-value pairs: ", m)
	fmt.Println("value 1: ", m["1"])

	delete(m, "2")
	fmt.Println("key-value pairs: ", m)

	for key, value := range m {
		fmt.Printf("Key: %s Value: %s \n", key, value)
	}
}
