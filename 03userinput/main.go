package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcome := "Welcome to user input"
	fmt.Println(welcome)

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter a message:")

	//comma ok || err err syntax

	input, _ := reader.ReadString('\n')
	fmt.Println("Thanks for the message,", input)

}
