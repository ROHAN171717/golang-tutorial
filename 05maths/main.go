package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
)

func main() {
	fmt.Println("Welcome to maths in golang");

	var x int = 2;
	var y int = 3;

	fmt.Println("Addition: ", x + y);

	//random number
	//rand.Seed(time.Now().UnixNano)
	//fmt.Println((rand.Intn(5)+1))

	//random from crypto
	randomNum, _ := rand.Int(rand.Reader, big.NewInt(5))
	fmt.Println(randomNum)

}