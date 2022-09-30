package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Welcome to Slices")

	var fruitList = []string{"Mango", "Apple"}

	fmt.Printf("Type of fruitlist is %T \n", fruitList)

	fruitList = append(fruitList, "Banana", "Toamto")
	fmt.Println(fruitList)

	fruitList = append(fruitList[:3])
	fmt.Println(fruitList)

	highScores := make([]int, 0)
	//default allocation of memory
	//if  nothing is initialised in slice , the default value is 0 just like arrays
	// highScores[0] = 10
	// highScores[1] = 4
	// highScores[2] = 5
	// highScores[3] = 8
	//highScores[4] = 15
	//panic: runtime error: index out of range [4] with length 4

	//append re-allocates the memory and the new values are accomodated
	highScores = append(highScores, 15, 60, 25)

	fmt.Println(highScores)

	//sort pkg
	fmt.Println(sort.IntsAreSorted(highScores))
	sort.Ints(highScores)
	fmt.Println(highScores)
	fmt.Println(sort.IntsAreSorted(highScores))

	//how to remove a value from slices based on index
	var courses = []string{"react.js", "javascript", "swift", "python", "ruby"}
	fmt.Println(courses)
	index := 2
	courses = append(courses[:index], courses[index+1:]...)
	fmt.Println(courses)
	// learn about ... syntax later

}
