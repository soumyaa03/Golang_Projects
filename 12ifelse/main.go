package main

import "fmt"

func main() {

	fmt.Println("If else in Golang")

	loginCount := 10
	var result string
	if loginCount < 10 {
		result = "Regular User"
	} else if loginCount > 10 {
		result = "Watch out"
	} else {
		result = "Exactly 10"
	}
	fmt.Println(result)

	if num := 3; num < 10 {
		fmt.Println("Number is less than 10")
	} else {
		fmt.Println("Number is more than 10")
	}

}
