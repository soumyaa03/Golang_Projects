package main

import "fmt"

func main() {

	fmt.Println("Maps in Goalng")

	languages := make(map[string]string)

	languages["JS"] = "Javascript"
	languages["RB"] = "Ruby"
	languages["PY"] = "Python"

	fmt.Println("list of languages : ", languages)
	fmt.Println("JS shorts for :", languages["JS"])

	//delete can be used for slices as well
	delete(languages, "RB")
	fmt.Println("list of languages :", languages)

	//loops are interesting in golang
	for key, value := range languages {
		fmt.Printf("For Key %v, value is %v \n", key, value)
	}

}
