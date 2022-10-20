package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	fmt.Println("welcome to files")
	content := "This need to go into a file - Golang@dev.in"

	file, err := os.Create("./myGolangfile.txt")
	if err != nil {
		panic(err)
	}

	length, err := io.WriteString(file, content)
	if err != nil {
		panic(err)
	}
	fmt.Println("length is : ", length)
	defer file.Close()
	readFile("./myGolangfile.txt")

}
func readFile(filname string) {
	databyte, err := ioutil.ReadFile(filname)
	if err != nil {
		panic(err)
	}
	fmt.Println("Text data inside the file is :", databyte)

	fmt.Println("Text data inside the file in String format is : ", string(databyte))
}
