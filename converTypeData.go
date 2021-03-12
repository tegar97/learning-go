package main

import "fmt"

func main() {
	var nilai32 int32 = 32768
	var nilai64 int64 = int64(nilai32)
	var nilai17 int16 = int16(nilai32)

	fmt.Println(nilai32)
	fmt.Println(nilai64)
	fmt.Println(nilai17)

	var name = "tegar"
	var e byte = name[0]
	var eString string = string(e)
	println(e)
	println(eString)

}
