package main

import "fmt"

func main() {
	defer fmt.Println("World 1")
	defer fmt.Println("World 2")
	defer fmt.Println("World 3")
	fmt.Println("Hello")
	myDefer()
	fmt.Println("")

}
func myDefer() {
	for i := 0; i < 5; i++ {
		defer fmt.Print(i)

	}

}
