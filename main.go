package main

import "fmt"

func print(str string) {
	fmt.Println("Variable address: \t", &str)
	fmt.Println("Variable value: \t", str)
}

func printPointer(str *string) {
	fmt.Println("Pointer address: \t", str)
	fmt.Println("Pointer value: \t\t", *str)
}

func main() {
	a := 25 // int
	b := &a // *int
	// c := &b // *(*int)

	// Could have written:
	var c *(*int) = &b
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)

	// TEST:

	myName := "MadMax"
	myPointer := &myName
	print(myName)
	printPointer(myPointer)

}
