package main 

import "fmt"

const LoginToken string = "itsokletsexplore"
//if we initialize a variable with capital letter then that variable publicly accessible
 

func main () {
	// var username string = "aarav"
	// var isGood bool = true
	// var smallVal uint8 = 255
	// var value int = 3456
	// var smallFloat float32 = 255.797979897987930
	var smallFloat float64 = 255.797979897987930
	fmt.Println(smallFloat )
	fmt.Printf("variable is of type : %T \n", smallFloat )

	// default values and some aliases
	var anotherVariable int //default value = 0
	fmt.Println(anotherVariable)
	fmt.Printf("variable is of type : %T \n", anotherVariable)


	//implicit type

	var website = "aaravshukla.me"
	fmt.Println(website)

	//no var style 

	numberOfUser := 30000.0
	fmt.Println(numberOfUser)

	fmt.Println(LoginToken)
	fmt.Printf("variable type of this public variable: %T", LoginToken)
}
