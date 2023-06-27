package main
import (
	"fmt"
	"net/http"
)

func main () {
	// go greeter("Hello")
	// greeter("World")

	websiteList := [] string {
		"https://lco.dev",
		"https://go.dev",
		"https://google.com",
		"https://fb.com",
		"https://github.com",
	}

	for _ , web := range websiteList {
		getStatusCode(web)
	}
} 



// func greeter (s string) {
// 	for i:= 0; i < 6 ;  i++{
// 		time.Sleep(5 * time.Millisecond)
// 		fmt.Println(s)
// 	}
// }


func getStatusCode (endpoint string){
	result, err := http.Get(endpoint)
	if err != nil{
		fmt.Println("endpoint error")
		
	}

	fmt.Printf(" %d status code for %s,", result.StatusCode, endpoint)
} 