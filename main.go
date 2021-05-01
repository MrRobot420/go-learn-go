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

	// Slices and arrays:

	// mySlice := make(<type>, <length>, <capacity>)
	mySlice := make([]string, 3)

	fmt.Printf("Length is: %d \nCapacity is: %d\n", len(mySlice), cap(mySlice))
	fmt.Println(mySlice)

	mySlice2 := []string{"apple", "orange", "peach"}
	fmt.Printf("Length is: %d \nCapacity is: %d\n", len(mySlice2), cap(mySlice2))
	fmt.Println(mySlice2)

	mySlice3 := []int{1, 2, 3, 4, 5, 6, 7}
	mySlice3 = append(mySlice3, 8)

	fmt.Printf("Length is: %d \nCapacity is: %d\n", len(mySlice3), cap(mySlice3))

	fmt.Println(mySlice3)

	mySlice4 := make([]int, 7, 12)
	fmt.Printf("Length is: %d \nCapacity is: %d\n", len(mySlice4), cap(mySlice4))

	for i := 0; i < 7; i++ {
		print("Adding: " + fmt.Sprint(i+1))
		mySlice4[i] = i + 1
	}
	fmt.Println(mySlice4)
	mySlice4 = append(mySlice4, 8)
	mySlice4 = append(mySlice4, 9)
	mySlice4 = append(mySlice4, 10)
	mySlice4 = append(mySlice4, 11)
	mySlice4 = append(mySlice4, 12)
	mySlice4 = append(mySlice4, 13)

	fmt.Printf("Length is: %d \nCapacity is: %d\n", len(mySlice4), cap(mySlice4))

	fmt.Println(mySlice4)

	initialSlice := []int{1, 2, 3, 4, 5, 6, 7}
	newSlice := []int{87, 56, 78}

	initialSlice = append(initialSlice, newSlice...)
	fmt.Println(initialSlice)
}
