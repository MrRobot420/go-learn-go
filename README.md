# The GO Lang

![alt text](https://yuriktech.com/images/Golang.png)


## Variables:

### Intro: 
Variables are used to store values in memory.
A bit is the smallest unit of data store in a computer.

Integers can be signed or unsigned.

Unsigned means its a positive number.
```go
uint8    // unsigned 8-bit integers (0 to 255)
uint16   // unsigned 16-bit integers (0 to 65535)
uint32   // unsigned 32-bit integers (0 to 4294967295)
uint64   // unsigned 64-bit integers (0 to A really big number)
```

Signed negative or positive.
```go
int8    // signed 8-bit integers (-128 to 127)
int16   // signed 16-bit integers (-32.768 to 32.767)
int32   // signed 32-bit integers (-2.147.483.648 to 2.147.483.647)
int64   // signed 64-bit integers (A really big negative number to a large number)
```

If you do not want to provide 8, 16, 32 or 64 to your int or uint, GO infers it for you.
```go
uint    // unsigned, either 32 or 64
int     // signed, either 32 or 64
```

Floats
```go
float32    // 32-bit floating-point numbers (1.24)
float64    // 64-bit floating-point numbers
```

Aliases
```go
byte        // alias for uint8
rune        // alias for int32
```

Booleans and Strings.
```go
bool        // true or false
string      // "Strings look like this and are surrounded by (double) quotes."
```

---

## How do we create a variable?

Explicit Declaration
We have to DECLARE with - var | variableName | type

Creates a location in memory to hold the declared TYPE

Location called greeting and hold a string

```go
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
```

---

## Pointers

& has ONE meaning.
The & symbol means: THE ADDRESS OF the variable it is next to.

Example: 
If a := 25 and we set b := &a,
the value of B is the memory address of A.
And its type is a pointer to an int - *int

```go
a := 25
b := &a
```

\* has TWO meanings
1: The * symbol next to a VARIABLE means: Get the value of the variable that this
pointer is pointing to aka DEREFERENCING.

Example: 
If we set n := *b
the value of n is 25. Its the value stored at the momory address
that b is pointing to.

```go
n := *n
```

2: The * symbol next to a TYPE means the variable being declared
is a pointer and it points to an address holding the type followed by the *.

Example: var *string myName
myName is a variable which holds the memory address of a string variable.

```go
var myName = "MadMax"
var myPointer *string
```


```go
func main() {
	a := 25 // int
	b := &a // *int
	c := &b // *(*int)

	// Could have written:
	// var c *(*int) = &b
	fmt.Println(*c)
}
```

---

## Control statements

#### 1. If / else statements:
```go
func main() {
    x := 5
    if x < 10 {
        fmt.Println("x is less than 10.")
    } else {
        fmt.Println("x is greater than or equals 10.")
    }
    fmt.Println("This will run no matter what.")
}
```

GO is block scoped.
However you can define a variable in the if part that will also be visible to else:
```go
func main() {
    x := 5
    if a:= x -3; x < 10 {
        fmt.Println("x is less than 10.")
    } else {
        fmt.Println(a)  // will work
    }
}
```

For loops:
```go
func main() {
    for i := 0; i < 5; i++ {
        fmt.Println(i)
    }
}
```

For loops for lists:
```go
func main() {
    names := []string{"Sam", "Tom", "Joe"}
    for i := 0; i < len(names); i++ {
        fmt.Println(names[i])
    }
}
```

"While" loops in GO are just for loops in another format.
```go
func main() {
    i := 0
    for i < 5 {
        fmt.Println(i)
        i += 1
    }
}
```

Endless loops in GO:
```go
func main() {
    i := 0
    for {
        fmt.Println(i)
        i += 1
        if i > 5 {
            break
        }
    }
}
```

For loops for a range of objects
```go
func main() {
    names := []string{"Sam", "Tom", "Joe"}
    for k, v := range names {
        fmt.Println(k, v)
    }
}
```

Switch statements:
```go
import (
    "fmt" 
    "time"
)

func main() {
    switch time.Now().Weekday() {
    case time.Saturday:
        fmt.Println("It's Saturday.")
    case time.Sunday:
        fmt.Println("It's Sunday.")
    default:
        fmt.Println("Today is a weekday.")
    }
}
```


IMPORTANT: Go BREAKS the switch statement, if the case was found! No fallthrough like in other languages where you have to "break" the switch explicitly.
```go
func main() {
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
```

---

### Slices and Arrays

An Array has a fixed length.
Slices are of flexible length.
Slices are just like Arrays, but are only a reference to an array (internally).

In order to make a slice, you need to call ```make()```. \
Make takes up to 3 params. \
The type, the length and the max size (capacity).
```go
func main() {
    // mySlice := make(<type>, <length>, <capacity>)
    mySlice := make([]string, 3, 8)

    fmt.Printf("Length is: %d \nCapacity is: %d\n", len(mySlice), cap(mySlice))
    fmt.Println(mySlice)
}
```

Alternate way of initializing a slice:
```go
func main() {
    mySlice := []string{"apple", "orange", "peach"}
    fmt.Printf("Length is: %d \nCapacity is: %d\n", len(mySlice), cap(mySlice))
    fmt.Println(mySlice)

    // print a slice of the slice ^^:
    fmt.Println(mySlice[1:4])   // [1:4] means: from index 1 up to but not including index 4

    fmt.Println(mySlice[:4]) // starting at 0 up to index 3

    fmt.Println(mySlice[1:]) // starting at 1 up to the last value (included!)
}
```

Append values to a slice:
```go
func main() {
    mySlice := []int{1, 2, 3, 4, 5, 6, 7}
    mySlice = append(mySlice, 8)

    fmt.Printf("Length is: %d \nCapacity is: %d\n", len(mySlice), cap(mySlice))

    fmt.Println(mySlice)
}
```

NOTICE: the capacity of the slice doubles every time that the length of the slice is higher than the previous capacity! \
Here the cap was 7. By adding the 8th value, the cap got updated to 14.

When using the ```make()``` function, you can prevent this.

You can also append values of a second slice to another slice.
```go
func main() {
    initialSlice := []int{1, 2, 3, 4, 5, 6, 7}
    newSlice := []int{87, 56, 78}

    initialSlice = append(initialSlice, newSlice...)    // Notice the dots. "elipsis"
    fmt.Println(initialSlice)
}
```

© Esoteric Tech @https://www.youtube.com/channel/UCg_k6DwjEs1wj4DngmxDSGQ