package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main () {
	fmt.Println("welcome to golang files")
	content := "this needs to go in a files"

	file, err := os.Create("./myfile.txt")

	checkNilErrror(err)

	length , err := io.WriteString(file, content)
	checkNilErrror(err)


	fmt.Println("length is ", length)

	defer file.Close()

	readFile("./myfile.txt")


}


func readFile (filename string) {
	dataByte, err := ioutil.ReadFile(filename)

	checkNilErrror(err);

	fmt.Println("text data inside the file is \n",string(dataByte))
}


func checkNilErrror (err error){

	if err != nil{
		panic(err)
	}
}
