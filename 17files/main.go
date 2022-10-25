package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	content := "This is my content"
	file, err := os.Create("./myFile.txt")
	checkNilErr(err)
	length, err := io.WriteString(file, content)
	checkNilErr(err)
	defer file.Close()
	fmt.Println("The length is", length)
	readFile("./myFile.txt")
}

func readFile(filename string) {
	dataByte, err := ioutil.ReadFile(filename)
	checkNilErr(err)
	fmt.Println("DATA INSIDE THE FILE IS:\n", string(dataByte))
}

func checkNilErr(err error) {
	if err != nil {
		panic(err)
	}
}
