package main

import "fmt"

func main() {
	var name = "tegar"

	switch name {
	case "tegar":
		fmt.Println("Hello Tegar")
	case "JOKO":
		fmt.Println("Hello Joko")
	default:
		fmt.Println("Kenalan")
	}

}
