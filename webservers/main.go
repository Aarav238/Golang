package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const url = "https://lco.dev"
func main () {
	fmt.Println("hello webservers")

	res, err := http.Get(url)

	if err != nil {
		panic(err)
	}

	fmt.Printf("response is of type : %T\n" , res)
	defer res.Body.Close()

	databytes,err := ioutil.ReadAll(res.Body)
	
	if err != nil {
		panic(err)
	}

	html := string(databytes)
	fmt.Println("the res body is: ", html  )
}