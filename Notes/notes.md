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
