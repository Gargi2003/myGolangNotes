package main

import "fmt"

func main() {
	var fruitlist [5]string
	//alternate method
	// var fruitlist = [5]string{"Apple", "Mango", "Banana", "Orange", "Kiwi"}

	fruitlist[0] = "Apple"
	fruitlist[1] = "Mango"
	fruitlist[3] = "Banana"
	fruitlist[4] = "peach"
	fmt.Println("fruitlist is: ", fruitlist)
	fmt.Println("fruitlist is: ", len(fruitlist))
}
