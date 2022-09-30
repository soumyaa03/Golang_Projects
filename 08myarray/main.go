package main

import "fmt"

func main() {
	fmt.Println("welcome to Array in Golang")
	var fruitList [4]string
	fruitList[0] = "Apple"
	fruitList[1] = "Tomato"
	fruitList[3] = "Mango"

	fmt.Println("Fruitlist is  : ", fruitList)
	fmt.Println("Length of fp is : ", len(fruitList))

	var vegList = [3]string{"potato", "beans", "mushroom"}
	fmt.Println("veg list is", len(vegList))

}
