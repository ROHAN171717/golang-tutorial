package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("### Slices in Golang ###")

	var xSlice = []string{"a", "b", "c"}
	fmt.Printf("xSlice: %T\n", xSlice)

	xSlice = append(xSlice, "d", "e")
	fmt.Println("xSlice: ", xSlice)

	xSlice = append(xSlice[:4])
	fmt.Println("xSlice: ", xSlice)

	z := make([]int, 4)

	z[0] = 17
	z[1] = 178
	z[2] = 174
	z[3] = 172

	z = append(z, 15, 18)

	fmt.Println("z: ", z)
	fmt.Println(sort.IntsAreSorted(z))
	sort.Ints(z)
	fmt.Println(sort.IntsAreSorted(z))
	fmt.Println(z)

	//how to remove a value from a slice based on index

	var w = []string{"a", "b", "c", "d", "e"}
	fmt.Println(w)
	var index int = 2
	w = append(w[:index], w[index+1:]...)
	fmt.Println(w)

}
