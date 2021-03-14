package main

import "fmt"

func random() interface{} {
	return "Ups"
}
func main() {
	var result interface{} = random()
	var resultString string = result.(string)
	fmt.Println(resultString)

}
