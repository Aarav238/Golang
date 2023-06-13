package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main (){

	fmt.Println("welcome to our app")
	fmt.Println("please rate it")

	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	fmt.Println("thanksfor rating,", input)

	numRating, err := strconv.ParseFloat(strings.TrimSpace( input) , 64)

	if err != nil {
		fmt.Println("err", err)
	
	} else{
		fmt.Println(numRating + 1)
	}



}