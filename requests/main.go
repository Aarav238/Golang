package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func main () {
	fmt.Println("hello world")

	PerformGetRequest()
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
