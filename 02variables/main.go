package main

import "fmt"

const LoggedIn string = "Gargi Bandyopadhyay" //public

func main() {
	var userName string = "Gargi"
	fmt.Println(userName)
	fmt.Printf("variable is of type: %T \n", userName)

	var isbool bool = true
	fmt.Println(isbool)
	fmt.Printf("variable is of type: %T \n", isbool)

	var float32Value float32 = 32.22643827912873276
	fmt.Println(float32Value)
	fmt.Printf("variable is of type: %T \n", float32Value)

	var float64Value float64 = 32.22643827912873276
	fmt.Println(float64Value)
	fmt.Printf("variable is of type: %T \n", float64Value)

	var integerValue int = 322
	fmt.Println(integerValue)
	fmt.Printf("variable is of type: %T \n", integerValue)

	var unitint uint8 = 32
	fmt.Println(unitint)
	fmt.Printf("variable is of type: %T \n", unitint)

	//default values

	var undefinedInt int
	fmt.Println(undefinedInt)
	fmt.Printf("variable is of type: %T \n", undefinedInt)

	var undefinedString string
	fmt.Println(undefinedString)
	fmt.Printf("variable is of type: %T \n", undefinedString)

	var undefinedBool bool
	fmt.Println(undefinedBool)
	fmt.Printf("variable is of type: %T \n", undefinedBool)

	var undefinedFloat float32
	fmt.Println(undefinedFloat)
	fmt.Printf("variable is of type: %T \n", undefinedFloat)

	//implicit type

	var website = "gargibanerjee49@gmail.com"
	fmt.Println(website)
	fmt.Printf("type of website: %T \n", website)

	//no var style

	numberOfUser := 3000
	fmt.Println(numberOfUser)

	fmt.Println(LoggedIn)
}
