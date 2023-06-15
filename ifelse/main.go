package main 
import . "fmt"

func main (){
	Println("IfElse in golang:")
	var result string

	loginCount:= 23
	if loginCount >= 10{
		result = "Not good"
	} else{
		result = "Regulate User"
	}

	Println(result)

	if 9%2 != 0 {
		Println("number is odd")
	} else {
		Println("number is even")		
	}

	if num := 3; num < 10 {
		Println("num is less than 10")
	} else {
		Println("Num is not less than 10")
	}

}