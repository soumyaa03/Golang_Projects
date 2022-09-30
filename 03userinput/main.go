package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcome := "Welcome to user input"
	fmt.Println(welcome)

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter rating for our pizza")
	input, _ := reader.ReadString('\n') // try catch type - ( ,_ :- here '_' catches the error)
	fmt.Println("Thanks for rating : " + input)

}
