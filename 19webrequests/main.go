package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const url = "https://lco.dev"

func main() {
	fmt.Println("welcome to web requests")
	response, err := http.Get(url)

	if err != nil {
		panic(err)
	}
	fmt.Printf("Type of response is  %T\n ", response)
	defer response.Body.Close() // caller's responsibility to close the connection

	databytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}
	content := string(databytes)
	fmt.Println(content)

}
