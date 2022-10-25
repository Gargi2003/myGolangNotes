package main

import (
	"fmt"
	"sort"
)

func main() {
	var fruitlist = []string{}
	fruitlist = append(fruitlist, "Apple", "Mango", "Banana", "Orange")
	fmt.Println("This is my fruitlist", fruitlist, len(fruitlist))
	fmt.Printf("Type: %T", fruitlist)

	//remove values from slice
	var index int = 2
	fruitlist = append(fruitlist[:index-1], fruitlist[index:]...)
	fmt.Println(fruitlist)

	//sort
	var numList = []int{100, 234, 456, 342, 674, 567}
	sort.Ints(numList)
	fmt.Println(numList)
	fmt.Println(sort.IntsAreSorted(numList))
}
