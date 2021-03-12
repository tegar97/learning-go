package main

import "fmt"

// Array adalah tipe data yang berisikan kumpulan data yang tipe sama
// saat membuat array, kita perlu menentukan jumlah data yang bisa ditampung oleh array tersebut
// Daya tampung array tidak bisa bertambah setelah array dibuat
func main() {
	var names [3]string
	names[0] = "tegar"
	names[1] = "akmal"

	fmt.Println(names[0])

	var buah = [3]string{
		"manga",
		"durian",
		"jeruk",
	}
	fmt.Println(buah)

}
