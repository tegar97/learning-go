package main

import "fmt"

/*
* Function tidak hanya bisa kita disimpan didalam variabel sebagai value
* Namun juga bisa kita gunakan sebagai parameter untuk function lain
 */

type Filter func(string) string

func sayHelloWithFilter(name string, filter Filter) {
	fmt.Println("Hello", filter(name))
}

func spamFilter(name string) string {
	if name == "Anjing" {
		return "...."
	} else {
		return name
	}
}
func main() {
	sayHelloWithFilter("Anjing", spamFilter)
}
