package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Please enter a number")
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')

	converted, err := strconv.ParseFloat(strings.TrimSpace(input), 64)
	if err != nil {
		fmt.Println(err)
	} else {
		// fmt.Println("Thanks for the input", input)
		fmt.Println("Thanks for the input", converted+1)
	}
}
