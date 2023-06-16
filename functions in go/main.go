package main
import "fmt"

func main(){

	hello()
	result := add(5,6)
	fmt.Println(result)
	proadder ,_:= proAdd(45,55,5,5,56,4444,7,5,5,444)
	fmt.Println(proadder)

}

func add (a int , b int) int {
	result:= a + b
	return result
}

func proAdd(values ...int ) (int,string) {
	total := 0

	for _, val := range values {
		total += val
	}

	return total , "helllo world"
}

func hello () {
	fmt.Println("hello functions ")
}

