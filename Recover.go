package main

import "fmt"

/*
	Recover
		* Recover adalah function yang bisa kita gunakan untuk menangkap data panic
		* Dengan recover proses panic akan terhenti , sehingga program akan tetap berjalan

*/
func endApp2() {
	message := recover()
	if message != nil {
		fmt.Println("Error dengan message", message)
	}
	fmt.Println("Aplikasi selesai")
}

func runApp2(error bool) {
	defer endApp2()
	if error {
		panic("Aplikasi error")
	}
	fmt.Println("Aplikasi berjalan")
}
func main() {
	runApp2(true)
	fmt.Println("test")
}
