package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

func main() {
	formData()
}
func formData() {
	const myurl = "http://localhost:1111/postform"

	data := url.Values{}

	data.Add("Company", "Dell Technologies")
	data.Add("Location", "Banglore")
	data.Add("Position", "Software Engineer 1")

	response, err := http.PostForm(myurl, data)
	if err != nil {
		panic(err)
	}
	content, _ := ioutil.ReadAll(response.Body)
	fmt.Println("Content is:", string(content))
}
