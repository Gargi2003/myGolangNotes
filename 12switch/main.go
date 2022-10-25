package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	rand.Seed(time.Now().UnixNano()) // seeding is done to get random nos ever single time
	diceNumber := rand.Intn(6) + 1
	fmt.Printf("dice num is: %v\n", diceNumber)
	switch diceNumber {
	case 1:
		fmt.Println("One")
	case 2:
		fmt.Println("Two")
	case 3:
		fmt.Println("Three")
	case 4:
		fmt.Println("Four")
	case 5:
		fmt.Println("Five")
	case 6:
		fmt.Println("Six")
	}
	fmt.Println("")
}
