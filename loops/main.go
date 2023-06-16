package main

import (
	"fmt"
)

func main () {
	fmt.Println("hello loops")

	days:= [] string {"sunday","tueday", "friday","Saturday"}

	 fmt.Println(days)
	// for d:=0 ; d<len(days);d++{
	// 	fmt.Println(days[d])
	// }

	// for i:= range days {
	// 	fmt.Println(days[i])
	// }

	// for index , day := range days {
	// 	fmt.Printf("index is %v and value is %v\n", index , day)
	// }

	rougueValue := 1

	for rougueValue <= 10 {

		// if rougueValue == 5{
		// 	break
		// }
		// if rougueValue == 5 {
		// 	rougueValue++
		// 	continue
		// }

		if rougueValue == 4 {
			goto lco
		}

		fmt.Println("value is ", rougueValue)
		rougueValue++;
	}


	lco:
		fmt.Println("learning golang")
}