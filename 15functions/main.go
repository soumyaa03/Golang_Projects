package main

import "fmt"

func main() {
	fmt.Println("Welcome to functions in golang ")
	result := adder(10, 5)
	fmt.Println("result is :", result)
	proRes, protext := proAdder(1, 2, 3, 4, 5, 6)
	fmt.Println("pro result is:", proRes)
	fmt.Println("pro message is:", protext)

}
func adder(val1 int, val2 int) int {
	return val1 + val2
}
func proAdder(values ...int) (int, string) {
	total := 0
	for _, value := range values {
		total += value
	}
	return total, "function completed successfully"
}
