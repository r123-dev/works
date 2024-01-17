package main

import "fmt"

func main() {

	languages := make(map[string]string)

	languages["JS"] = "Javascript"
	languages["Rb"] = "Ruby"
	languages["PY"] = "Python"

	fmt.Println("List of all languages: ", languages)
    for key,value := range languages{
		fmt.Printf(" For key %v, value is %v\n", key, value)
	}

}