package main

import (
	"fmt"
	"time"
)

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

	x := 5
	if x < 10 {
		fmt.Println("x is less than 10.")
	} else {
		fmt.Println("x is greater than or equals 10.")
	}
	fmt.Println("This will run no matter what.")

	y := 5
	if a := y - 3; a > 3 {
		fmt.Println("y is less than 10.")
	} else {
		fmt.Println(a) // will work
	}

	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	names := []string{"Sam", "Tom", "Joe"}
	for i := 0; i < len(names); i++ {
		fmt.Println(names[i])
	}

	i := 0
	for i < 5 {
		fmt.Println(i)
		i += 1
	}

	i = 0
	for {
		fmt.Println(i)
		i += 1
		if i > 5 {
			break
		}
	}

	for k, v := range names {
		fmt.Println(k, v)
	}

	switch time.Now().Weekday() {
	case time.Saturday:
		fmt.Println("It's Saturday.")
	case time.Sunday:
		fmt.Println("It's Sunday.")
	default:
		fmt.Println("Today is a weekday.")
	}

	// Example of auto-break in switch cases
	switch 2 {
	case 1:
		fmt.Println("1")
	case 2:
		fmt.Println("2")
	case 3:
		fmt.Println("3")
	}

	// Example of fallthrough in switch cases
	switch 2 {
	case 1:
		fmt.Println("1")
		fallthrough
	case 2:
		fmt.Println("2")
		fallthrough
	case 3:
		fmt.Println("3")
	}
}
