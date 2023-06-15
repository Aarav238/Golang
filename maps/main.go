package main

import (
	"fmt"
)

func main(){
	fmt.Println("maps in golang")

	langauges := make(map[string]string)
	langauges["JS"] = "javascript"
	langauges["RB"] = "Ruby"
	langauges["PY"] = "Python"

	fmt.Println("list of all language", langauges)
	fmt.Println("JS shorts for :", langauges["JS"])

	delete(langauges, "RB")
	fmt.Println("list of all language", langauges)


	//loops

	for key , value := range langauges {
		fmt.Printf("For key %v , value is %v \n", key , value)

	}
		
}


