package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

func main() {
	fmt.Println("Welcome to webrequests")
	// PerformGetRequest()
	// PerformPostJsonRequest()
	PerformPostFormRequest()
}

func PerformGetRequest() {
	const myurl = "http://localhost:8000/get"

	response, err := http.Get(myurl)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()
	fmt.Println("Status code is : ", response.StatusCode)
	fmt.Println("Content length : ", response.ContentLength)

	var responseString strings.Builder
	content, _ := ioutil.ReadAll(response.Body)
	responseString.Write(content)
	// bytecount,_ := responseString.Write(content)

	// fmt.Println("Byte Count is: ", bytecount)
	fmt.Println(responseString.String())
	// fmt.Println(string(content))
}

func PerformPostJsonRequest() {
	const myurl = "http://localhost:8000/post"

	requestBody := strings.NewReader(`
	{
		"coursename" : "golang",
		"price" : 0,
		"platform" : "learnCodeOnline.in"
	}
	`)
	response, err := http.Post(myurl, "application/json", requestBody)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()

	content, _ := ioutil.ReadAll(response.Body)
	fmt.Println(string(content))
}

func PerformPostFormRequest() {
	const myurl = "http://localhost:8000/postform"
	//formdata

	data := url.Values{}
	data.Add("firstname", "Soumya")
	data.Add("lastname", "Hota")
	data.Add("email", "Soumya@go.dev")

	response, err := http.PostForm(myurl, data)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()
	content, _ := ioutil.ReadAll(response.Body)
	fmt.Println(string(content))
}
