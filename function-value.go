package main

import "fmt"

// Function adalah first class citizen
// Function juga merupakan tipe data dan bisa disimpan  dalam variabel

func getGoodBye(name string) string {
	return "Good Bye " + name
}
func main() {
	goodBye := getGoodBye
	fmt.Println(goodBye("tegar"))
}
