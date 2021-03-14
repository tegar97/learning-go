package main

import "fmt"

/*
	Nil
	* Biasanya di dalam bahasa pemograman lain,object yang belum dinisialisai maka secara otomatis nilainya
      adalah null atau nill
	* Berbeda dengan Go-Lang , di Go-Lang saat kita buat variabel dengan tipe data tertentu ,maka secara
      otomatis akan dibuatkan defaultnya
    * Namun di Go-Land ada data nil,yaitu data kosong
	* Nill sendiri hanya bisa digunakan di beberapa tipe data seperti interface,function,map,slice,pointer dan
	  channel
*/

func NewMap(name string) map[string]string {
	if name == "" {
		return nil
	} else {
		return map[string]string{
			"name": name,
		}
	}
}

func main() {
	var person map[string]string = NewMap("TEGAR")
	fmt.Println(person)

}
