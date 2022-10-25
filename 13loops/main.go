package main

import "fmt"

func main() {
	days := []string{"Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday", "Sunday"}

	//loop 1
	for i := 0; i < len(days); i++ {
		fmt.Println(days[i])
	}

	//loop2
	for i := range days {
		fmt.Println(days[i])
	}
	// loop3 for each
	for index, day := range days {
		fmt.Printf("index is %v and day is %v\n", index, day)
	}

	//loop4 while
	randomValue := 1
	for randomValue < 10 {
		if randomValue == 3 {
			goto goesto
		}
		if randomValue == 5 {
			break
		} else if randomValue == 6 {
			randomValue++
			continue
		}
		fmt.Println("value is:", randomValue)
		randomValue++
	}

goesto:
	fmt.Println("Going to the label")
}
