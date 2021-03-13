package main

/*
										Variadic Function
	* Parameter yang berada di posisi terakhir , memiliki kemampuan dijadikan sebuah varags
	* Varargs artinya datanya bisa menerima lebih dari satu input , atau anggap asaja semacam Array.
	* Apa bedanya dengan parameter biasa dengan tipe data Array ?
		* Jika parameter tipe Array , kita wajib membuat array terlebih dahulu sebelum mengirimkan ke function
		* Jika parameter menggunakan varargs , kita bisa langsung mengirim datanya , jika lebih dari satu,cukup
           gunakan tanda koma

*/

func sumAll(numbers ...int) int {
	total := 0
	for _, number := range numbers {
		total += number
	}
	return total
}

func main() {
	total := sumAll(20, 30, 40, 50)
	println(total)
}
