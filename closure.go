package main

import "fmt"

/*
* Closure adalah kemampuan sebuah function berinterasi dengan data-data disekitarnya dalam scope yang sama

 */
func main() {
	name := "tegar"
	counter := 0

	increment := func() {
		name := "Budi"
		fmt.Println("increment")
		counter++
		fmt.Println(name)
	}

	increment()

	fmt.Println(counter)
	fmt.Println(name)
}
