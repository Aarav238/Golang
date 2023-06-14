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
