package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func main() {
	postCall()
}

func postCall() {
	const url = "http://localhost:1111/post"

	data := strings.NewReader(`
		{
			"Company":"Dell Technologies",
			"Location":"Banglore",
			"Position":"Software Engineer 1"

		}
	`)

	response, err := http.Post(url, "application/json", data)
	if err != nil {
		panic(err)
	}
	content, _ := ioutil.ReadAll(response.Body)
	fmt.Println("Content is: ", string(content))
	defer response.Body.Close()
}
