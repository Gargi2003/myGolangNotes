package main

import "fmt"

func main() {
	gargi := User{"Gargi", "gargi@gmail.com", true, 24}
	fmt.Printf("Details are %+v\n", gargi)
	fmt.Printf("name is %v and email is %v\n", gargi.Name, gargi.Email)
	gargi.GetStatus()
	gargi.NewEmail()
	fmt.Println(gargi.Email) //original ale of the email gargi@gmail.com doesnt change even after setting it to newEmail. This indicated that when we pass arguments to a method, a copy of these variables are passed and notthe actual values. Hence the need of pointers was felt.
}

// no inheritence no super no parent
type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

func (u User) GetStatus() {
	fmt.Println("Status of the user is", u.Status)
}
func (u User) NewEmail() {
	u.Email = "newEmail@gmail.com"
	fmt.Println("New email is:", u.Email)
}
