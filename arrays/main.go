package main 
import "fmt"

func main () {
	fmt.Println("Array")

	var fruitList [4]string
	fruitList[0]="Apple"
	fruitList[2]="Mango"
	fruitList[3]= "Orange"

	fmt.Println(fruitList)
	fmt.Println(len(fruitList))
	var veglist = [3]string {"patato" , "tamato" , "brinjal"}
	fmt.Println(veglist)
	fmt.Println(len(veglist))

}