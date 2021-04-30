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

### How do we create a variable?

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





© Esoteric Tech @https://www.youtube.com/channel/UCg_k6DwjEs1wj4DngmxDSGQ