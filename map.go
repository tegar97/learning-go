package main

import "fmt"

func main() {
	person := map[string]string{
		"name":   "tegar",
		"adress": "bandung",
	}
	person["title"] = "Programmer"
	fmt.Println(person)
}
