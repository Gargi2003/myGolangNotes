package main

import "fmt"

func main() {
	languages := make(map[string]string)
	languages["JS"] = "Javascript"
	languages["PY"] = "Python"
	languages["TS"] = "Typescript"
	languages["RB"] = "Ruby"
	fmt.Println(languages)
	//delete
	delete(languages, "JS")
	fmt.Println(languages)

	//iterating through map
	for k, v := range languages {
		fmt.Printf("values for %v are %v\n", k, v)
	}

}
