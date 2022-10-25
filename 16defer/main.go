/*
In Golang, the defer keyword is used to delay the execution of a function or a statement until the nearby function returns.
In simple words, defer will move the execution of the statement to the very end inside a function.
If multiple lines are defered then the order of reexecuting the defered statements would be in LIFO fashion. This is the reason the output of the below code is Hello-> World -> Two ->One
*/

package main

import "fmt"

func main() {
	//example 1
	defer fmt.Println("One")
	defer fmt.Println("Two")
	defer fmt.Println("World")
	fmt.Println("Hello")

	//example 2
	greet()

	//example 3
	myDefer()
}

//example 2
func greet() {
	defer fmt.Println("Hey I am inside greet")
	fmt.Println("Defered")
}

//example 3
func myDefer() {
	for i := 0; i < 5; i++ {
		defer fmt.Print(i)
	}
}
