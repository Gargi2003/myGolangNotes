package main

import "fmt"

func main() {
	result, message := adder(1, 2, 3, 5, 6)
	fmt.Println("result is:", result, message)
}
func adder(values ...int) (int, string) {
	total := 0
	for _, value := range values {
		total += value
	}
	return total, "This is a message"
}
