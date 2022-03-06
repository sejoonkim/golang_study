package main

import "fmt"

func main() {
	fmt.Println("Welcome to functions in golang")
	greeter()
	// greeterTwo()
	result := adder(3, 5)
	fmt.Println("Result is :", result)

	proRes, myMessage := proAdder(2, 3, 5, 6, 6, 4)
	fmt.Println("Result is :", proRes)
	fmt.Println("Message is :", myMessage)
}

// func greeterTwo() {
// 	fmt.Println("Another method")
// }

func adder(valOne int, valTwo int) int {
	return valOne + valTwo
}

func proAdder(values ...int) (int, string) {
	total := 0

	for _, val := range values {
		total += val
	}

	return total, "HI"
}

func greeter() {
	fmt.Println("Namastey from golang")
}
