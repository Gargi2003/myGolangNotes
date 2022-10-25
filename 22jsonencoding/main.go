package main

import (
	"encoding/json"
	"fmt"
)

type Course struct {
	Name     string   `json:"courseName"` //changes the key Name to courseName in final json
	Price    int      `json:"price"`
	Platform string   `json:"website"`
	Password string   `json:"-"`              // hides this entity from the final json
	Tags     []string `json:"tags,omitempty"` //discards nil from the final json
}

func main() {
	encodeJson()
}

func encodeJson() {
	courses := []Course{
		{"Golang Bootcamp", 120, "Learnonlinecoding.com", "abc123", []string{"webdev", "go"}},
		{"Javascript Bootcamp", 100, "Learnonlinecoding.com", "abcd123", []string{"webdev", "js"}},
		{"Python Bootcamp", 320, "Learnonlinecoding.com", "aabc123", nil},
	}

	finalJson, _ := json.MarshalIndent(courses, "", "\t")
	fmt.Printf("%s", finalJson)
}
