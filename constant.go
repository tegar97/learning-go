package main

import "fmt"

// Constant adlah variabel yang nilainya tidak bisa diubah lagi setelah pertama kali di beri nilai
// Cara pembuatan constant mirip dengan variabel yang membedakan hanya kata kunci yang digunakan adalah const,bukan var.
// Saat pembuatan constant, kita wajib langsung menginisialisasikan datanya
func main() {
	const firstname string = "tegar"
	const lastname = "akmal"

	fmt.Println(firstname," ",lastname)

}
