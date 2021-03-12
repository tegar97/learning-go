package main

// Type Declartions adalah kemampuan membuat ulang tipe data baru dari tpe data yang sudah ada
// Type Declartions biasanya digunakan untuk membuat alias terhadap tipe data yang sudah ada
// Dengan Tujuan agar lebih mudah dimengerti
func main() {
	type noktp string
	type Married bool
	var noKtpTegar noktp = "323012412501251262161261"
	var marriedStatus Married = false
	println(noKtpTegar)
	println(marriedStatus)
}
