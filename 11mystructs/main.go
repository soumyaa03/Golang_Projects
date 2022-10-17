package main

import "fmt"

func main() {

	fmt.Println("Structs in Golang")
	// no inheritance in golang; No super or parent

	soumya := User{"Soumya", "soumya@go.dev", true, 21}
	fmt.Println(soumya)
	//%v for value , %+v for more detailation of struct
	fmt.Printf("soumya's details are : %+v\n", soumya)
	fmt.Printf("Name is %v and email is %v.", soumya.Name, soumya.Email)

}

type User struct {
	//first letter capital means "public"
	Name   string
	Email  string
	Status bool
	Age    int
}
