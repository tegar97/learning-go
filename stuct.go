package main

import "fmt"

/*
	Struct
		* Struct adalah sebuah template data yang  digunakan untuk menggabungkan  nol atau lebih tipe data
          lainya dalam satu kesatuan
		* Struct biasanya representasi data dalam program aplikasi yang kita buat
		* Data di struct disimpan dalam field
		* Sederhanya struct adalah kumpulan dari field



*/

type Customers struct {
	Name, Adress string
	Age          int
}

func main() {
	var tegar Customers
	tegar.Name = "Tegar Akmal"
	tegar.Adress = "Indonesia"
	tegar.Age = 30
	fmt.Println(tegar)

	budi := Customers{
		Name:   "Budi",
		Adress: "Bandung",
		Age:    20,
	}
	fmt.Println(budi)

}
