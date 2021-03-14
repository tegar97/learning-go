package main

import "fmt"

/*
	Pointer di function
	* Saat kita membuat parameter di function,secara default adalah pass by value , artinya data akan di copy
	  lalu dikirim ke function tersebut
	* Oleh karena itu , jika kita mengubah data didalam function,data yang aslinya tidak akan pernah berubah
	* Hal ini membuat variabel  menjadi aman, karena tidak akan bisa diubah
	* Namun kadang kita ingin membuat function yang bisa mengubah data asli paremeter tersebut
	* Untuk melakukan ini , kita juga bisa menggunakan pointer di function
	* Untuk menjadikan sebuah parameter sebagai pointer, kita bisa menggunakan operator di * parameternya
*/

func ChangeCountryToIndonesia(adress *Adress) {
	adress.Country = "Argetina"
}
func main() {
	adress := Adress{"bandung", "jawa barat", "indonesia"}
	ChangeCountryToIndonesia(&adress)
	fmt.Println(adress)
}
