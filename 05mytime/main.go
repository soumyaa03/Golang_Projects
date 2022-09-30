package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Welcome to time study of golang")
	presentTime := time.Now()
	fmt.Println(presentTime)
	//defalt format for printing date
	//"01-02-2006" will give us today's date
	//"15:04:05" will give current time
	//"Monday" will return which weekly day it is
	fmt.Println(presentTime.Format("01-02-2006 15:04:05 Monday"))
	createdDate := time.Date(2022, time.January, 12, 23, 0, 0, 0, time.UTC)
	fmt.Println(createdDate)
	fmt.Println(createdDate.Format("01-02-2006 Monday"))
}
