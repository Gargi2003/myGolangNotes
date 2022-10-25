package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("This is today's date")
	currentDate := time.Now()
	fmt.Println(currentDate)

	formattedDate := currentDate.Format("2006-01-02 Monday")
	fmt.Println("This is the formatted date:", formattedDate)

	createdDate := time.Date(2022, time.December, 20, 15, 4, 5, 6, time.UTC)
	fmt.Println("This is created date", createdDate.Format("2006-01-02 Monday"))
}
