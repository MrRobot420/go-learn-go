package main

import "fmt"

var greeting string

func main() {
	// Store a value in our variable aka INITIALIZATION
	greeting = "What it is"
	fmt.Println(greeting)

	// IMPLICIT Declaration
	// Declare AND initialize a variable using :=
	a := 30
	fmt.Println(a)
	year := "2021"
	myString := "Happy New Year! it's "
	myBool := 5 > 8

	// We can also use var with implicit declaration
	var myFloat float32 = 5.8 + 2.4

	fmt.Println(myString + year)
	fmt.Println(myBool)
	fmt.Println(myFloat)
}
