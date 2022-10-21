package main

import (
	"encoding/json"
	"fmt"
)

type course struct {
	Name     string `json:"coursename"`
	Price    int
	Platform string   `json:"website"`
	Password string   `json:"-"`
	Tags     []string `json:"tags,omitempty"`
}

func main() {
	fmt.Println("Welcome to Json")
	EncodeJson()

}

func EncodeJson() {

	myCourses := []course{
		{"Flutter Bootcamp", 300, "learnGo.in", "abc123", []string{"flutter", "dart"}},
		{"GoLang Bootcamp", 300, "learnGo.in", "abc456", []string{"node-express", "js"}},
		{"ReactJS Bootcamp", 300, "learnGo.in", "abc789", nil},
	}
	//package this data as JSON data
	finalJson, err := json.MarshalIndent(myCourses, "", "\t")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", finalJson)

}
