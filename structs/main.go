package main 
import "fmt"

func main () {
	fmt.Println("hello structs")
	//no inheritance in golang, No super or parent

Aarav:= User{"Aarav", "Aarav@go.dev", true, 16}

fmt.Println(Aarav)
fmt.Printf("Aarav details are: %+v \n", Aarav)
fmt.Printf("His name is: %v \n", Aarav.Name)


}

type User struct {
	Name  string
	Email  string
	Status bool
	Age    int
}