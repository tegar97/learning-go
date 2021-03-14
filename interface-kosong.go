package main

import "fmt"

/*
	Interface Kosong
	* Gol-lang bukanlah bahasa pemrograman yang berorientasi objek
	* Biasanya dalam pemograman berorientasi object,ada satu data parent di puncak yang biasa dianggap sebagai semua
      implemtasi data yang ada di bahasa pemrograman tersebut
	* Contoh di java ada java.lang.Object
	* Untuk menangani kasus seperti ini , di Go-Lang kita bisa menggunakan interface kosong
	* interface kosong adalah interface yang tidak memilki deklarasi method satupun,hal ini membuat secara
	  otomatis semua tipe data akan menjadi implmentasinya

*/
func Ups(i int) interface{} {
	if i == 1 {
		return 1
	} else if i == 2 {
		return true
	} else {
		return "Ups"
	}
}
func main() {
	var data interface{} = Ups(1)
	fmt.Println(data)
}
