package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const url = "https://lco.dev"

func main() {
	response, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Response is of type %T\n", response)
	defer response.Body.Close() //callers responsibility to close the connection
	databytes, err := ioutil.ReadAll(response.Body)
	fmt.Println("The data is", string(databytes))
	if err != nil {
		panic(err)
	}
}
