# Packages and Imports in Go

## Packages

- Go code is organized into packages.
- Each package has a unique name and contains related code files.
- Package declaration is done using the `package` keyword.
- Example: `package main` declares the main package for an executable program.
```go
package main

import "fmt"

func main() {
    fmt.Println("Hello, Go!")
}
```

## Imports

- Packages are imported to use code from other packages.
- The `import` keyword is used for importing packages.
- Import statements are typically placed at the top of the Go file.
- Example: `import "fmt"` imports the "fmt" package for I/O operations.
```go
import "fmt"

func main() {
    fmt.Println("Hello, Go!")
}

```

## Using Imported Packages

- Imported package identifiers are accessed using the package name followed by a dot (`.`).
- Example: `fmt.Println("Hello, Go!")` uses the `Println` function from the "fmt" package.
```go
import "fmt"

func main() {
    fmt.Println("Hello, Go!")
}

```

## Aliasing Imports

- Packages can be imported with an alias to avoid naming conflicts.
- The `as` keyword is used to specify the alias.
- Example: `import fm "fmt"` imports the "fmt" package with the alias `fm`.
```go
import fm "fmt"

func main() {
    fm.Println("Hello, Go!")
}

```
## Importing Specific Identifiers

- Specific identifiers can be imported from a package using the dot (`.`) notation.
- This allows direct use of imported identifiers without the package name.
- Example: `import . "fmt"` imports all the exported identifiers from the "fmt" package.
```go
import . "fmt"

func main() {
    Println("Hello, Go!")
}

```

Remember, proper package organization and importing the necessary packages help create modular and readable Go code.



---
# Variables in Go

## Declaration and Initialization

- Variables in Go are declared using the `var` keyword followed by the variable name and type.
- Variables can be initialized at the time of declaration or later.
- The `:=` short variable declaration syntax can be used for short and concise variable declaration and initialization.

Example:
```go
var age int // Variable declaration
age = 25 // Variable initialization

name := "John" // Short variable declaration and initialization
```
## Variable Types

- Go is a statically-typed language, which means variables have a fixed type that cannot be changed.
- Common variable types include `int` for integers, `float64` for floating-point numbers, `string` for text, `bool` for boolean values, and more.
- You can also create custom types using the `type` keyword.

**Example:**
```go
var count int
var pi float64
var message string
var isActive bool
```

## Constants

- Constants are similar to variables but their values cannot be changed once assigned.
- Constants are declared using the `const` keyword.
- Constants must be assigned a value at the time of declaration.

**Example:**
```go
const pi = 3.14
const daysInWeek = 7
```
## Scope and Visibility

- Variables in Go have a scope, which determines where the variable is accessible.
- A variable declared within a function has a local scope and is only accessible within that function.
- A variable declared outside of any function, at the package level, has a global scope and can be accessed by all functions within that package.

**Example:**
```go
package main

import "fmt"

var globalVariable = "Global" // Global variable

func main() {
    var localVariable = "Local" // Local variable
    fmt.Println(localVariable)
    fmt.Println(globalVariable)
}
```

---
# User Input in Go

## Reading User Input

- To read user input in Go, you can use the `bufio` package along with `os.Stdin` for standard input.
- First, you need to import the necessary packages: `bufio` and `os`.
- Create a new `Scanner` using `bufio.NewScanner(os.Stdin)` to read input from the user.
- Use the `Scanner`'s `Scan()` method to read the input until a newline character is encountered.
- The user input can be accessed using the `Text()` method of the `Scanner`.

Example:
```go
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcome := "welcome to user input"
	fmt.Println(welcome)

	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Enter the rating for goland:")

	//comma okay syntax or error ok syntax

	input, _ := reader.ReadString('\n')

	fmt.Println(" thanks for rating ,", input)
	fmt.Printf("type of this rating is %T", input)
}

```

# Time Conversion Study in Go

## Getting Current Time

- To get the current time in Go, you can use the `time.Now()` function.
- Import the `time` package to access time-related functionality.
- Call `time.Now()` to obtain the current time.
- The returned value is of type `Time`.

Example:
```go
presentTime := time.Now()
fmt.Println(presentTime)
```

## Formatting Time
- You can format a `Time` value into a specific layout using the `Format()` method.
- The `Format()` method accepts a layout string that defines the desired time format.
- The layout string is based on specific format placeholders like "01" for month, "02" for day, "2006" for year, and so on.Different placeholders are separated by specific separator characters, such as "-" or ":".
- Import the `time` package to access time-related functionality.

Example: 
```go
fmt.Println(presentTime.Format("01-01-2006 15:04:05 Monday "))

```
above should be followed strictly.
## Creating a Custom Date

- You can create a custom `Time` value representing a specific date and time using the `time.Date()` function.
- Specify the year, month, day, hour, minute, second, and time zone information to create the `Time` value.
- The time zone information is optional; you can omit it or use `time.UTC` for UTC time.
- Import the `time` package to access time-related functionality.

**Example:**
```go
createdDate := time.Date(2020, time.January, 10, 23, 23, 0, 0, time.UTC)
fmt.Println(createdDate)
fmt.Println(createdDate.Format("01-02-2006 Monday "))
```
In the example above, we create a Time value called createdDate representing January 10, 2020, at 23:23 in the UTC time zone. We then print the createdDate value, which outputs the date and time in the default format. Next, we format the createdDate value using the layout string "01-02-2006 Monday", which prints the date in the format "01-02-2006" (month-day-year) followed by the day of the week.

Remember to import the necessary packages (fmt and time), use the appropriate functions and methods from the time package, and format the time using the desired layout string.

# Type Conversion in Go

## Implicit Type Conversion (Type Coercion)

- Go performs automatic type conversion when assigning a value of one type to a variable of another compatible type.
- This is known as implicit type conversion or type coercion.
- Implicit type conversion is allowed when there is no loss of precision or potential data loss.

Example:
```go
var num int = 10
var result float64 = num // Implicit conversion from int to float64
```
## Explicit Type Conversion (Type Casting)

- Explicit type conversion, also known as type casting, is used to convert a value of one type to another explicitly.
- Type casting is done using the syntax `Type(value)` where `Type` represents the target type and `value` is the value to be converted.
- Explicit type conversion is necessary when there is potential loss of precision or when converting between incompatible types.

**Example:**
```go
var num float64 = 3.14
var result int = int(num) // Explicit conversion from float64 to int
```
In the example above, we have a variable `num` of type `float64` that stores the value `3.14`. We want to convert this floating-point value to an integer. We use explicit type conversion by using the `int()` function and passing `num` as the argument. The result is stored in the `result` variable, which is now of type `int`.

Explicit type conversion should be used when there is a need to convert between types that are not automatically compatible. It allows you to control the conversion process and handle any potential loss of precision or data loss.

Remember to be cautious when performing explicit type conversion and ensure that it is appropriate for the specific conversion you need to perform.

## String Conversion

- Converting other types to strings is done using the `strconv` package or using the `fmt.Sprintf` function.
- The `strconv.Itoa` function is used to convert an integer to a string.
- The `strconv.FormatFloat` function is used to convert a float to a string.
- The `fmt.Sprintf` function can be used to convert values of various types to formatted strings.

**Example:**
```go
import (
    "strconv"
    "fmt"
)

var num int = 10
var strNum string = strconv.Itoa(num) // Convert int to string

var pi float64 = 3.14
var strPi string = strconv.FormatFloat(pi, 'f', 2, 64) // Convert float to string with 2 decimal places

var message string = fmt.Sprintf("The value of pi is %.2f", pi) // Convert float to formatted string
```

In the example above, we have demonstrated different methods for converting other types to strings:

- To convert an `int` value to a string, we use the `strconv.Itoa()` function. In the example, we convert the `num` variable of type `int` (with a value of `10`) to a string using `strconv.Itoa(num)`.
- To convert a `float64` value to a string, we use the `strconv.FormatFloat()` function. In the example, we convert the `pi` variable of type `float64` (with a value of `3.14`) to a string with 2 decimal places using `strconv.FormatFloat(pi, 'f', 2, 64)`.
- To convert values of various types to formatted strings, we can use the `fmt.Sprintf()` function. In the example, we create a formatted string using the value of `pi` and the `%f` verb to represent the float value with 2 decimal places.

String conversion is useful when you need to represent values of different types as strings, such as for printing or formatting purposes.

Remember to import the necessary packages (`strconv` and `fmt`) to use the provided functions for string conversion.










---
Go tool can run file directly , without VM
Executables are different OS

What

system apps to web apps - Cloud
Already in production

This languages don't have some features but those features are don't needed.
Golang don't have classes they have struct.

Lexer - ensure that you are following the grammer of language



case insensitive; almost
Everything is a type

GOOS="windows" go build

memory allocation and deallocation happens automatically

new()                    make()
allocatie memory but   allocate memory and INIT
no INIT
you will get memory    you will get a memory 
address                    address

zeroed storage               non-zeroed storage


GC happens automatically
out of scope or nil

reference means &
