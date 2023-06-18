package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func main () {
	fmt.Println("hello world")

	//PerformGetRequest()
	PerformPostJSONRequest()
}


func PerformGetRequest(){
	const myurl = "http://localhost:8000/get"

	res,err :=  http.Get(myurl)

	if err!=nil{
		panic(err)
	}
	defer res.Body.Close()
	fmt.Println("Status code:",res.StatusCode)
	fmt.Println("Content length:" , res.ContentLength)

	var responseString strings.Builder
	content,err := ioutil.ReadAll(res.Body)
	byteCount, _ := responseString.Write(content)

	fmt.Println(responseString.String())

	fmt.Println(byteCount)
	if err!=nil{
		panic(err)
	}

	//fmt.Println(string(content))
	

}


func PerformPostJSONRequest (){
	myurl := "http://localhost:8000/post"

	//fake json

	requestBody:= strings.NewReader(`
		{
			"name":"Aarav",
			"email":"aarav@gmail.com",
			"phone":"898989898"
		}
	`)

	res ,err := http.Post(myurl, "application/json", requestBody)
	
	
	if err!=nil{
		panic(err)
		
	}
	defer res.Body.Close()
	content,err  := ioutil.ReadAll(res.Body)
	
	if err!=nil{
		panic(err)
	}
	fmt.Println(string(content))
}
