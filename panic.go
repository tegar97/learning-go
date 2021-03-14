package main

import "fmt"

/*
	Panic
		* Panic function adalah function yang bisa kita gunakan untuk menghentikan program
		* Panic function biasanya dipanggil ketika terjadi error pada saat program kita berjalan
		* saat panic function dipanggil,program akan terhenti,namun defer function tetap akan di eksekusi


*/

func endApp() {
	fmt.Println("Aplikasi selesai")
}

func runApp(error bool) {
	defer endApp()
	if error {
		panic("Aplikasi error")
	}
	fmt.Println("Aplikasi berjalan")
}
func main() {
	runApp(true)

}
