package main

import (
	"fmt"
	"sort"
)

func main(){

	var fruitList = []string{"hello","tata","bye","bye"}

	fmt.Printf("Type of fruitlisti is %T\n", fruitList)
	fruitList  = append(fruitList, "see", "you")
	fmt.Println(fruitList)

	fruitList = append(fruitList[1:4])
	fruitList = append(fruitList[1:])
	
	
	fmt.Println(fruitList)


	highScore := make([]int,4)

	highScore[0]  = 98345
	highScore[1]  = 2374
	highScore[2]  = 48485
	highScore[3]  = 8989

	highScore = append(highScore, 666,777,77776)

	fmt.Println(len(highScore))
	fmt.Println(sort.IntsAreSorted(highScore))
	sort.Ints(highScore) 
	fmt.Println(highScore)
	fmt.Println(sort.IntsAreSorted(highScore))

	}