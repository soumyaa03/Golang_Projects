package main

import "fmt"

func main() {
	fmt.Println("Welcome to a aclass on pointers")
	// pointer if empty , then it returns 0
	// var ptr *int
	// fmt.Println("Value of pointer is ", ptr)

	myNumber := 23
	// reference means '&'
	var ptr = &myNumber
	fmt.Println("value of the pointer is : ", ptr)
	fmt.Println("value of the pointer is : ", *ptr)

	*ptr = *ptr + 2
	fmt.Println("new value is : ", myNumber)
}
