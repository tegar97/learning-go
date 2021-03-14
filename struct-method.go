package main

import "fmt"

/*
	Struct
	* Struct adalah tipe ddata seperti tipe data lainya , dia bisa digunakan sebagai parameter untuk function
	* Namun jika kita ingin menambahkan method ke dalam stucts , sehingga seakan akan sebuah sruct memiliki function
	* Method adalah function
*/

type Customers2 struct {
	Name, Adress string
	Age          int
}

func (Customers Customers2) sayHi(name string) {
	fmt.Println("Hello", name, "my name is", Customers.Name)
}
func main() {
	var tegar Customers2
	tegar.Name = "Tegar Akmal"
	tegar.Adress = "Indonesia"
	tegar.Age = 30

	tegar.sayHi("tegar")
	fmt.Println(tegar)

	budi := Customers2{
		Name:   "Budi",
		Adress: "Bandung",
		Age:    20,
	}
	fmt.Println(budi)

}
