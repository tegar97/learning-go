package main

import "fmt"

// Named Return Values
// Biasanya saat kita memberi tahu bahwa semuah function mengembalikan value , maka kita  hanya
// mendeklarasikan tipe data return value di function
// Namun kita juga bisa membuat variabel secara langsung di tipe data return function nya

func getFullName() (firstName, middleName, lastName string) {
	firstName = "tegar"
	middleName = "akmal"
	lastName = "2"
	return

}

func main() {
	firstName, middleName, lastName := getFullName()
	fmt.Println(firstName)
	fmt.Println(middleName)
	fmt.Println(lastName)

}
