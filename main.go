package main

import "fmt"

func main() {
	a := 25 // int
	b := &a // *int
	// c := &b // *(*int)

	// Could have written:
	var c *(*int) = &b
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
}
