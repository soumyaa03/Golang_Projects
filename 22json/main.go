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
	DecodeJson()

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

func DecodeJson() {

	jsonDataFromWeb := []byte(`
	{
		
        
			"coursename": "Flutter Bootcamp",
			"Price": 300,
			"website": "learnGo.in",
			"tags": ["flutter","dart"]
	
	}`)
	var devCourse course
	checkValid := json.Valid(jsonDataFromWeb)
	if checkValid {
		fmt.Println("JSON was valid")
		json.Unmarshal(jsonDataFromWeb, &devCourse)
		//%#v for printing interfaces
		fmt.Printf("%#v\n", devCourse)
	} else {
		fmt.Println("JSON was not valid")
	}

	//some cases where we just want to add key value pairs
	var onlineData map[string]interface{}
	json.Unmarshal(jsonDataFromWeb, &onlineData)
	fmt.Printf("%#v\n", onlineData)

	for k, v := range onlineData {
		fmt.Printf("key is %v and value is %v and type is : %T\n", k, v, v)
	}
}
