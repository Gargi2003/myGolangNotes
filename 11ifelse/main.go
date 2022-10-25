package main

import "fmt"

func main() {
	if num := 3; num < 10 {
		fmt.Println("num<10")
	} else {
		fmt.Println("num>=10")
	}

	if 9%2 == 0 {
		fmt.Println("even")
	} else {
		fmt.Println("odd")
	}

}
