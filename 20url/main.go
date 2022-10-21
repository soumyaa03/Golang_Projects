package main

import (
	"fmt"
	"net/url"
)

const myurl string = "https://lco.dev:3000/learn?coursename=reactjs&paymentid=ghdjhfs74"

func main() {
	fmt.Println("Welcome to Urls")
	fmt.Println(myurl)

	//parsing
	result, _ := url.Parse(myurl)

	// fmt.Println(result.Scheme)
	// fmt.Println(result.Host)
	// fmt.Println(result.Path)
	// fmt.Println(result.Port())
	// fmt.Println(result.RawQuery)

	qparams := result.Query()
	fmt.Printf("the type of query parameters are : %T\n", qparams)

	//parameters are in key-value pairs
	// fmt.Println(qparams["coursename"])
	// fmt.Println(qparams["paymentid"])

	for _, value := range qparams {
		fmt.Println("Param is : ", value)
	}

	//always use a referense - "&"
	partsofUrl := &url.URL{
		Scheme:  "https",
		Host:    "lco.dev",
		Path:    "/tutcss",
		RawPath: "user=soumya",
	}

	anotherUrl := partsofUrl.String()
	fmt.Println(anotherUrl)

}
