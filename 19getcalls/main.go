package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	getRequest()
}

func getRequest() {
	const url = "http://localhost:1111"

	response, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	content, err := ioutil.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println("Content: ", string(content))

	/*

		alternate way to convert byte to string

		var responseString strings.Builder
		byteData, _ := responseString.Write(content)
		fmt.Println("Byte Data and length", responseString.String(), byteData)

	*/

	fmt.Println("Content Length: ", response.ContentLength)
	fmt.Println("Request Status: ", response.StatusCode)
	defer response.Body.Close()
}
