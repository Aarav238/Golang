package main

import (
	"fmt"
	"net/url"
)

const myurl string = "https://lco.dev:3000/learn?coursename=reactjs&paymentid=kldjl494"

 

func main() {
	fmt.Println("welcome to url handler")
	fmt.Println(myurl)

	//parsing

	result,err := url.Parse(myurl)

	if err != nil {
		panic(err)
	}

	fmt.Println(result.Scheme)
	fmt.Println(result.Host)
	fmt.Println(result.Path)
	fmt.Println(result.Port())
	fmt.Println(result.RawQuery)


	queryParams := result.Query()
	fmt.Printf("the type of query is %T\n", queryParams)

	fmt.Println("the type of query is :", queryParams)
	fmt.Println(queryParams["coursename"])

	for _ , value := range queryParams {

		fmt.Println("param is ", value)
	}

	partOfUrl := &url.URL{
		Scheme: "https",
		Host: "lco.dev",
		Path: "/tutcss",
		RawPath: "user=aarav",
	}

	anotherUrl := partOfUrl.String()
	fmt.Println(anotherUrl)	
	
	
	

	
	
	
	

}