package main

import "fmt"

/*
	* Sebelumnya setiap membuat function , kita akan selalu memberikan sebuah nama pada function
      tersebut
	* Namun kadang ada kalanya lebih mudah  membuat function secara langsung di variabel atau parameter
	  tanpa harus membuat funtion terlebih dahulu
    * Hal tersebut dinamakan anonymous function atau function tanpa nama

*/

type BlackList func(string) bool

func registerUser(name string, blacklist BlackList) {
	if blacklist(name) == true {
		fmt.Println("Blocked", name)
	} else {
		fmt.Println("Welcome", name)
	}
}

func main() {
	blacklist := func(name string) bool {
		return name == "admin"
	}

	registerUser("admin", blacklist)

	registerUser("root", blacklist)
}
