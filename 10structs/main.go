package main

import "fmt"

func main() {
	gargi := User{"Gargi", "gargi@gmail.com", true, 24}
	fmt.Println(gargi)
	fmt.Printf("Details are %+v\n", gargi)
	fmt.Printf("name is %v and email is %v", gargi.Name, gargi.Email)
}

//no inheritence no super no parent
type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}
