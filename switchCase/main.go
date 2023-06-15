package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main () {
	fmt.Println("hello")


	rand.Seed(time.Now().UnixNano())
	diceNumber := rand.Intn(6) + 1

	fmt.Println("value of dice is:", diceNumber)

	switch diceNumber {
	
			case 1 : 
			fmt.Println("Move to one spot")
			case 2 : 
			fmt.Println("Move to two spot")
			fallthrough

			case 3 : 
			fmt.Println("Move to three spot")

			case 4 : 
			fmt.Println("Move to four spot")

			case 5 : 
			fmt.Println("Move to five spot")

			case 6 : 
			fmt.Println("Move to six spot and roll dice again")

			default : 
			fmt.Println("what was that")
	

	}

}