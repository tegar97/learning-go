package main

import "fmt"

/*
	Defer
	 	* Defer function adalah function yang bisa kita jadwalkan untuk di eksekusi setelah sebuah function
			selesai di eksekusi
		* Defer function akan selalu diekseksi walaupun terjadi error di function yang di eksekusi

*/
func loggin() {
	fmt.Println("Selesai memanggil function")
}

func runApplication(value int) {
	defer loggin()
	fmt.Println("Run Application")
	result := 10 / value
	fmt.Println("resukt", result)
}

func main() {
	runApplication(0)
}
