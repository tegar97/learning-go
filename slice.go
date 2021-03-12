package main

import "fmt"

func main() {
	var month = [...]string{
		"januari",
		"februari",
		"maret",
		"april",
		"mei",
		"juni",
		"juli",
		"agustus",
		"september",
		"oktober",
		"november",
		"desember",
	}
	fmt.Println(month)
	var slice1 = month[3:7]

	fmt.Println(cap(slice1))
	//var newMonths = append(slice1, "desember baru")
	//println(newMonths)
}
