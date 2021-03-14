package main

import "fmt"

/*
	Pointer
	 A.Pass By Value
      * secara default di Go-lang semua variabel di passing by value , bukan by reference
      * Artinya , jika kita mengirim sebua variabel ke dalam function , method atau variabel lain,
		sebenarnya yang dikirim adalah duplikasi valuenya

	* Pointer adalah kemampuan membuat reference ke lokasi data di memory yang sama , tanpa menduplikasi data
	  yang sudah ada
     * Sederhananya , dengan kemampuan pointer , kita bisa membuat pass by reference


*/

type Adress struct {
	City, Province, Country string
}

func main() {
	var adrress1 = Adress{
		City:     "Bandung",
		Province: "Jawa barat",
		Country:  "Indonesia",
	}
	var adress2 = &adrress1
	var adress3 = &adrress1
	adress2.City = "JAKARTA"

	*adress2 = Adress{
		City:     "Bandung",
		Province: "Jawa timur",
		Country:  "Indonesia",
	}
	fmt.Println(adrress1)
	fmt.Println(adress2)
	fmt.Println(adress3)
}
