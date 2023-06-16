package main

import "fmt"
func main (){
	defer fmt.Println("World")
	defer fmt.Println("Hello")
	defer fmt.Println("India")
	defer fmt.Println("Australia")
	defer fmt.Println("England")
	defer fmt.Println("New Zealand")
	fmt.Println("I")
	fmt.Println("Am")
	fmt.Println("Aarav")
	fmt.Println("Shukla")
	myDefer()
	

}


func myDefer(){

	for i:=0; i < 5 ; i++{
		defer fmt.Println(i)
	}
}