package main

import "fmt"

func main() {
	var counter = 1
	for counter <= 100000000 {
		fmt.Println("Perulangan ke ", counter)
		counter++
	}
}
