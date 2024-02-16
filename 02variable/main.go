package main

import "fmt";

const token string = "asdfg";

func main() {
	var var1 string = "Asd";
	fmt.Println(var1);
	fmt.Printf("Variable is of type: %T \n", var1);

	var var2 bool = false;
	fmt.Println(var2);
	fmt.Printf("Variable is of type: %T \n", var2);

	var var3 uint8 = 255;
	fmt.Println(var3);
	fmt.Printf("Variable is of type: %T \n", var3);

	var var4 float64 = 256.555;
	fmt.Println(var4);
	fmt.Printf("Variable is of type: %T \n", var4);

	var var5 int = 255;
	fmt.Println(var5);
	fmt.Printf("Variable is of type: %T \n", var5);

	var var6 float32 = 255.55;
	fmt.Println(var6);
	fmt.Printf("Variable is of type: %T \n", var6);

	var var7 string;
	var7 = "ASD";
	fmt.Println(var7);
	fmt.Printf("Variable is of type: %T \n", var7);

	var8 := "ASD";
	fmt.Println(var8);
	fmt.Printf("Variable is of type: %T \n", var8);

}