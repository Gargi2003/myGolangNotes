package main

import "fmt"

func main() {
	number := 23
	var ptr = &number
	fmt.Println("actual pointer/memory is ", ptr)
	fmt.Println("actual value inside memory is ", *ptr)
	*ptr += 2
	fmt.Println("new value", number)
}
