package main

import "fmt"


// variable adalah tempat untuk penyimpanan data
// variabel digunakan agar kita bisa mengakses data yang sama dimanapun kita mau
// Di Go lang variable hanya bisa menyimpan tipe data yang sama , jika kita ingin menyimpan data yang berbeda
// -beda jenis.kita harus membuat beberapa variabel
// Untuk membuat variabel, kita bisa menggunakan kata kunci var,lalu diikuti dengan nama variabel dan tipe datanya

func main() {
	var name string
	name = "tegar"
	fmt.Println(name)
	fmt.Println(len(name))

	var umur int8
	umur  = 17
	fmt.Println(umur)
}
