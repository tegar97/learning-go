package main

import "fmt"

/*
	* Interface adalah tipe data Abstact , dia tidak memiliki implementasi langsung
	* Sebuah interface berisikan definisi-definisi method
	* biasanya interface digunakan sebagai kontrak

	implementasi interace
	* setiap tipe data yang sesuai dengan kontrak interface, secara otomatis dianggap sebagai interaface
      tersebut
	* Sehingga kita tidak perlu mengimplementasikan interface secara manual
	* Hal ini agak berbeda dengan bahasa pemograman lain yang ketika membuat interface, kita harus menyebutkan secara eksplisit akan menggunakan interface mana

*/

type HasName interface {
	GetName() string
}

func sayHello2(HasName HasName) {
	fmt.Println("Hello", HasName.GetName())
}

type Person struct {
	name string
}

func (person Person) GetName() string {
	return person.name
}

func main() {
	var tegar Person

	tegar.name = "tegar"

	sayHello2(tegar)

}
