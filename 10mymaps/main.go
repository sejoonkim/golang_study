package main

import "fmt"

func main() {
	fmt.Println("Maps in golnag")

	languages := make(map[string]string)
	languages["JS"] = "Javascript"
	languages["RB"] = "Ruby"
	languages["PY"] = "Python"

	fmt.Println("List of all languages:", languages)
	fmt.Println("List of all languages:", languages["JS"])
	fmt.Println("List of all languages:", languages["RB"])

	delete(languages, "RB")
	fmt.Println("List of all languages:", languages)

	// loops are interesting in golang
	for _, value := range languages {
		fmt.Printf("For key v, value is %v\n", value)
	}
}
