package main

import "fmt"

/*
	Recursive function adalah function yang memanggil function dirinya sendiri
	Kadang dalam pekerjaan , kita sering menemui kasus dimana menggunakan recursive function lebih
	mudah dibanding tidak menggunakan recurvie function
	Contoh kasus yang lebih mudah diselesaikan menggunakan recursive adalah factorial
*/

func factorialLoop(value int) int {
	result := 1
	for i := value; i > 0; i-- {
		result *= i
	}

	return result
}

/* ====== Recursive ====== */

func factorialRecursive(value int) int {
	if value == 1 {
		return 1
	} else {
		return value * factorialRecursive(value-1)
	}
}
func main() {
	loop := factorialLoop(5)
	loop2 := factorialRecursive(5)
	fmt.Println(loop)
	fmt.Println(loop2)

}
